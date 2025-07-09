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
	"log"
)

type ServiceInventoryRepository interface {
	GetByID(context context.Context, id string, selectFields *string) (*models.Service, error)
	Create(context context.Context, serviceCreate *models.ServiceCreate) (*mongo.InsertOneResult, error)
	GetAllPaginate(context context.Context, queryParams bson.M, selectFields *string, offset *int64, limit *int64) ([]*models.Service, *int64, error)
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

func (repo *MongoServiceInventoryRepository) GetAllPaginate(context context.Context, queryParams bson.M, selectFields *string, offset *int64, limit *int64) ([]*models.Service, *int64, error) {

	offset, limit = utils.ValidatePaginationParams(offset, limit)

	findOptions := &options.FindOptions{ // Find options
		Skip:  offset,
		Limit: limit,
	}

	// Fields Projection
	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		findOptions.SetProjection(fieldProjection)
	}

	// Get list of service orders

	records, err := repo.MongoCollection.Find(context, queryParams, findOptions)
	if err != nil {
		return nil, nil, err
	}

	retrieveServiceOrders := []*models.Service{}
	for records.Next(context) {
		var serviceOrder = models.Service{}
		if err := records.Decode(&serviceOrder); err != nil {
			log.Println("Error decoding document:", err)
			continue
		}
		retrieveServiceOrders = append(retrieveServiceOrders, &serviceOrder) // Append pointer
	}

	totalCount, err := repo.MongoCollection.CountDocuments(context, queryParams)

	if err != nil {
		return nil, nil, err
	}

	return retrieveServiceOrders, &totalCount, nil
}
