package repository

import (
	"context"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/utils"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

type ServiceInventoryRepository interface {
	GetByID(context context.Context, id string, selectFields *string) (*models.Service, error)
	Create(context context.Context, serviceCreate *models.ServiceCreate) (*models.Service, error)
	GetAllPaginate(context context.Context, httpRequest *http.Request, selectFields *string, offset *int64, limit *int64) ([]*models.Service, *int64, error)
}

type MongoServiceInventoryRepository struct {
	MongoCollection *mongo.Collection
	Logger          *core.Logger
}

func (repo *MongoServiceInventoryRepository) GetByID(context context.Context, id string, selectFields *string) (*models.Service, error) {
	//serviceId, err := primitive.ObjectIDFromHex(id)

	//if err != nil {
	//	return nil, errors.New("id is not valid")
	//}

	// Apply projection if set
	findOptions := options.FindOne()
	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		findOptions.SetProjection(fieldProjection)
	}

	filter := bson.D{{Key: "id", Value: id}}
	record := repo.MongoCollection.FindOne(context, filter, findOptions)

	retrieveService := models.Service{}
	err := record.Decode(&retrieveService)

	if err != nil {
		return nil, err
	}

	return &retrieveService, nil
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

func (repo *MongoServiceInventoryRepository) GetAllPaginate(context context.Context, httpRequest *http.Request, selectFields *string, offset *int64, limit *int64) ([]*models.Service, *int64, error) {

	offset, limit = utils.ValidatePaginationParams(offset, limit)
	fieldProjection := utils.GerFieldsProjection(selectFields)

	// Get filter or pipeline
	queryParams := httpRequest.URL.Query()
	depth := -1
	if deepVals, ok := queryParams["deep"]; ok && len(deepVals) > 0 {
		if d, err := strconv.Atoi(deepVals[0]); err == nil {
			depth = d
		}
		delete(queryParams, "deep")
	}

	//TODO:  This logic move to service layer
	if depth >= 0 { // pipeline mode
		filterOrPipeline, _ := utils.BuildTmfMongoFilter(queryParams, true)
		pipeline := filterOrPipeline.(mongo.Pipeline)
		pipeline = append(pipeline,
			bson.D{{Key: "$graphLookup", Value: bson.M{
				"from":             "serviceInventory",
				"startWith":        "$serviceRelationship.service.id",
				"connectFromField": "serviceRelationship.service.id",
				"connectToField":   "id",
				"as":               "relatedServices",
				"depthField":       "level",
				"maxDepth":         depth,
			}}},
		)
		// Add pagination stages
		if offset != nil && limit != nil {
			pipeline = append(pipeline,
				bson.D{{Key: "$skip", Value: *offset}},
				bson.D{{Key: "$limit", Value: *limit}},
			)
		}

		// Add projection if defined
		if len(fieldProjection) > 0 {

			pipeline = append(pipeline,
				bson.D{{Key: "$project", Value: fieldProjection}},
			)
		}

		cursor, err := repo.MongoCollection.Aggregate(context, pipeline)
		if err != nil {
			return nil, nil, err
		}

		var results []*models.Service
		//var rawResults []bson.M
		if err := cursor.All(context, &results); err != nil {
			return nil, nil, err
		}

		// For aggregate, total count isn't trivial – can omit or add $count stage separately if needed
		total := int64(len(results))
		return results, &total, nil
	} else {
		filterOrPipeline, _ := utils.BuildTmfMongoFilter(queryParams, false)
		findOptions := &options.FindOptions{
			Skip:  offset,
			Limit: limit,
		}

		if len(fieldProjection) > 0 {
			findOptions.SetProjection(fieldProjection)
		}

		cursor, err := repo.MongoCollection.Find(context, filterOrPipeline.(bson.M), findOptions)
		if err != nil {
			return nil, nil, err
		}

		var results []*models.Service
		if err := cursor.All(context, &results); err != nil {
			return nil, nil, err
		}

		totalCount, err := repo.MongoCollection.CountDocuments(context, filterOrPipeline.(bson.M))
		if err != nil {
			return nil, nil, err
		}

		return results, &totalCount, nil
	}
}
