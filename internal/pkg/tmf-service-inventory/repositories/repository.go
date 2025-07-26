package repository

import (
	"context"
	"errors"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type ServiceInventoryRepository interface {
	GetByID(context context.Context, id string, selectFields *string) (*models.Service, error)
	Create(context context.Context, serviceCreate *models.ServiceCreate) (*mongo.InsertOneResult, error)
	GetAllPaginate(context context.Context, httpRequest *http.Request, selectFields *string, offset *int64, limit *int64) ([]*models.Service, *int64, error)
}

type MongoServiceInventoryRepository struct {
	MongoCollection *mongo.Collection
	Logger          *core.Logger
}

func (repo *MongoServiceInventoryRepository) GetByID(context context.Context, id string, selectFields *string) (*models.Service, error) {
	serviceId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("id is not valid")
	}

	// Apply projection if set
	findOptions := options.FindOne()
	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		findOptions.SetProjection(fieldProjection)
	}

	filter := bson.D{{Key: "_id", Value: serviceId}}
	record := repo.MongoCollection.FindOne(context, filter, findOptions)

	retrieveService := models.Service{}
	err = record.Decode(&retrieveService)

	if err != nil {
		return nil, err
	}

	return &retrieveService, nil
}
func (repo *MongoServiceInventoryRepository) Create(context context.Context, serviceCreate *models.ServiceCreate) (*mongo.InsertOneResult, error) {

	insertResult, err := repo.MongoCollection.InsertOne(context, serviceCreate)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func (repo *MongoServiceInventoryRepository) GetAllPaginate(context context.Context, httpRequest *http.Request, selectFields *string, offset *int64, limit *int64) ([]*models.Service, *int64, error) {

	offset, limit = utils.ValidatePaginationParams(offset, limit)

	// Get filter or pipeline
	filterOrPipeline, isPipeline := utils.BuildTmfMongoFilter(httpRequest.URL.Query())

	// Fields Projection
	fieldProjection := utils.GerFieldsProjection(selectFields)
	if isPipeline {
		// pipeline mode
		pipeline := filterOrPipeline.(mongo.Pipeline)

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
		// regular find
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
