package repository

import (
	"context"
	"errors"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type ServiceOrderRepository interface {
	GetByID(context context.Context, id string, selectFields *string) (*models.ServiceOrder, error)
	Create(context context.Context, serviceCreate *models.ServiceOrder) (*models.ServiceOrder, error)
	Delete(context context.Context, id string) (bool, error)
	GetAllPaginate(context context.Context, httpRequest *http.Request, selectFields *string, offset *int64, limit *int64) ([]*models.ServiceOrder, *int64, error)
	Update(context context.Context, id string, serviceOrder interface{}) (bool, error)
}

type MongoServiceOrderRepository struct {
	MongoCollection *mongo.Collection
	MongoClient     *mongo.Client
	Logger          *core.Logger
}

func (repo *MongoServiceOrderRepository) GetByID(context context.Context, id string, selectFields *string) (*models.ServiceOrder, error) {

	// Apply projection if set
	findOptions := options.FindOne()
	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		findOptions.SetProjection(fieldProjection)
	}

	filter := bson.D{{Key: "id", Value: id}}
	record := repo.MongoCollection.FindOne(context, filter, findOptions)

	retrieveService := models.ServiceOrder{}
	err := record.Decode(&retrieveService)

	if err != nil {
		return nil, err
	}

	return &retrieveService, nil
}
func (repo *MongoServiceOrderRepository) Create(context context.Context, serviceOrderCreate *models.ServiceOrder) (*models.ServiceOrder, error) {

	_, err := repo.MongoCollection.InsertOne(context, serviceOrderCreate)

	if err != nil {
		return nil, err
	}

	return serviceOrderCreate, nil
}

func (repo *MongoServiceOrderRepository) Delete(context context.Context, id string) (bool, error) {

	filter := bson.D{{Key: "id", Value: id}}
	deleteRecord, err := repo.MongoCollection.DeleteOne(context, filter)

	if err != nil {
		return false, err
	}

	if deleteRecord.DeletedCount == 0 {
		return false, errors.New("Delete record with ID:" + id + " not success")
	}

	// filter document where serviceRelationship.service.id is deleted document and remove that from list
	//_, err = repo.MongoCollection.UpdateMany(
	//	context,
	//	bson.M{"serviceRelationship.service.id": id},
	//	bson.M{"$pull": bson.M{
	//		"serviceRelationship": bson.M{
	//			"service.id": id,
	//		},
	//	}},
	//)
	//
	//if err != nil {
	//	return false, err
	//}

	return true, nil
}

func (repo *MongoServiceOrderRepository) GetAllPaginate(context context.Context, httpRequest *http.Request, selectFields *string, offset *int64, limit *int64) ([]*models.ServiceOrder, *int64, error) {

	offset, limit = utils.ValidatePaginationParams(offset, limit)
	mongoPipeline := mongo.Pipeline{
		bson.D{{Key: "$skip", Value: *offset}},
		bson.D{{Key: "$limit", Value: *limit}},
	}

	// Add projection if defined
	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 {
		mongoPipeline = append(mongoPipeline,
			bson.D{{Key: "$project", Value: fieldProjection}},
		)
	}

	// Apply Filter
	mongoFilter := utils.BuildTmfMongoFilter(httpRequest.URL.Query())
	mongoPipeline = append(mongoPipeline,
		bson.D{{Key: "$match", Value: mongoFilter}},
	)

	cursor, err := repo.MongoCollection.Aggregate(context, mongoPipeline)
	if err != nil {
		return nil, nil, err
	}

	var results []*models.ServiceOrder
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

func (repo *MongoServiceOrderRepository) Update(context context.Context, id string, serviceOrder interface{}) (bool, error) {

	if _, ok := serviceOrder.(*models.ServiceOrderUpdate); !ok {
		if _, ok := serviceOrder.(*models.ServiceOrder); !ok {
			return false, errors.New("Invalid service order type")
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
		update := bson.M{"$set": serviceOrder}
		result := repo.MongoCollection.FindOneAndUpdate(context, filter, update)

		if err = result.Err(); err != nil {
			return nil, err
		}

		// Decode the result
		updatedServiceOrder := models.ServiceOrder{}
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
