package repository

import (
	"context"
	"errors"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/utils"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
	"time"
)

type ServiceInventoryRepository interface {
	GetByID(context context.Context, id string, selectFields *string, GraphLookupDepth *int64) (*models.Service, error)
	Create(context context.Context, serviceCreate *models.ServiceCreate) (*models.Service, error)
	Delete(context context.Context, id string) (bool, error)
	GetAllPaginate(context context.Context, httpRequest *http.Request, selectFields *string, offset *int64, limit *int64) ([]*models.Service, *int64, error)
	Update(context context.Context, id string, service interface{}) (bool, error)
}

type MongoServiceInventoryRepository struct {
	MongoCollection *mongo.Collection
	MongoClient     *mongo.Client
	Logger          *core.Logger
}

func (repo *MongoServiceInventoryRepository) GetByID(ctx context.Context, id string, selectFields *string, graphLookupDepth *int64) (*models.Service, error) {

	mongoPipeline := mongo.Pipeline{
		{{"$match", bson.D{{"id", id}}}},
		{{"$limit", 1}},
	}

	if graphLookupDepth != nil && *graphLookupDepth >= 0 {
		mongoPipeline = append(mongoPipeline,
			bson.D{{Key: "$graphLookup", Value: bson.M{
				"from":             repo.MongoCollection.Name(),
				"startWith":        "$serviceRelationship.service.id",
				"connectFromField": "serviceRelationship.service.id",
				"connectToField":   "id",
				"as":               "graphLookup",
				"depthField":       "graphLookupDepth",
				"maxDepth":         *graphLookupDepth,
			}}},
		)
	}

	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		mongoPipeline = append(mongoPipeline,
			bson.D{{Key: "$project", Value: fieldProjection}},
		)
	}
	record, err := repo.MongoCollection.Aggregate(ctx, mongoPipeline)
	defer func() {
		_ = record.Close(context.Background())
	}()
	if err != nil {
		return nil, err
	}

	retrieveService := models.Service{}
	if record.Next(ctx) {
		err = record.Decode(&retrieveService)
		if err != nil {
			return nil, err
		}

		return &retrieveService, nil

	}

	return nil, mongo.ErrNoDocuments
}
func (repo *MongoServiceInventoryRepository) Create(context context.Context, serviceCreate *models.ServiceCreate) (*models.Service, error) {

	service := models.Service{}
	err := copier.Copy(&service, serviceCreate)

	if err != nil {
		return nil, err
	}

	uid := uuid.New().String()
	service.ID = &uid

	createDate := strfmt.DateTime(time.Now().UTC())
	service.ServiceDate = &createDate

	_, err = repo.MongoCollection.InsertOne(context, service)

	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (repo *MongoServiceInventoryRepository) Delete(context context.Context, id string) (bool, error) {

	filter := bson.D{{Key: "id", Value: id}}
	deleteRecord, err := repo.MongoCollection.DeleteOne(context, filter)

	if err != nil {
		return false, err
	}

	if deleteRecord.DeletedCount == 0 {
		return false, errors.New("Delete record with ID:" + id + " not success")
	}

	// filter document where serviceRelationship.service.id is deleted and remove that from list
	_, err = repo.MongoCollection.UpdateMany(
		context,
		bson.M{"serviceRelationship.service.id": id},
		bson.M{"$pull": bson.M{
			"serviceRelationship": bson.M{
				"service.id": id,
			},
		}},
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *MongoServiceInventoryRepository) GetAllPaginate(context context.Context, httpRequest *http.Request, selectFields *string, offset *int64, limit *int64) ([]*models.Service, *int64, error) {

	offset, limit = utils.ValidatePaginationParams(offset, limit)
	mongoPipeline := mongo.Pipeline{
		bson.D{{Key: "$skip", Value: *offset}},
		bson.D{{Key: "$limit", Value: *limit}},
	}

	queryParams := httpRequest.URL.Query()
	graphLookupDepth := -1
	if deepVals, ok := queryParams["graphLookupDepth"]; ok && len(deepVals) > 0 {
		if d, err := strconv.Atoi(deepVals[0]); err == nil {
			graphLookupDepth = d
		}
		delete(queryParams, "graphLookupDepth")
	}

	if graphLookupDepth >= 0 {
		mongoPipeline = append(mongoPipeline,
			bson.D{{Key: "$graphLookup", Value: bson.M{
				"from":             repo.MongoCollection.Name(),
				"startWith":        "$serviceRelationship.service.id",
				"connectFromField": "serviceRelationship.service.id",
				"connectToField":   "id",
				"as":               "graphLookup",
				"depthField":       "graphLookupDepth",
				"maxDepth":         graphLookupDepth,
			}}},
		)
	}

	// Apply Filter
	mongoFilter := utils.BuildTmfMongoFilter(queryParams)
	mongoPipeline = append(mongoPipeline,
		bson.D{{Key: "$match", Value: mongoFilter}},
	)

	// Add projection if defined
	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 {
		mongoPipeline = append(mongoPipeline,
			bson.D{{Key: "$project", Value: fieldProjection}},
		)
	}

	cursor, err := repo.MongoCollection.Aggregate(context, mongoPipeline)
	if err != nil {
		return nil, nil, err
	}

	var results []*models.Service
	//var rawResults []bson.M
	if err := cursor.All(context, &results); err != nil {
		return nil, nil, err
	}

	total, err := repo.MongoCollection.CountDocuments(context, mongoFilter)
	if err != nil {
		return nil, nil, err
	}

	return results, &total, nil
}

func (repo *MongoServiceInventoryRepository) Update(context context.Context, id string, service interface{}) (bool, error) {

	if _, ok := service.(*models.ServiceUpdate); !ok {
		if _, ok := service.(*models.Service); !ok {
			return false, errors.New("Invalid service type")
		}
	}

	// Start a new session
	session, err := repo.MongoClient.StartSession()
	if err != nil {
		return false, err
	}
	defer session.EndSession(context)

	_, err = session.WithTransaction(context, func(sessCtx mongo.SessionContext) (interface{}, error) {
		filter := bson.M{"id": id}
		update := bson.M{"$set": service}
		result := repo.MongoCollection.FindOneAndUpdate(context, filter, update)

		if err = result.Err(); err != nil {
			return nil, err
		}

		// Decode the result
		updatedServiceOrder := models.Service{}
		if err = result.Decode(&updatedServiceOrder); err != nil {
			return nil, err
		}

		return true, nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
