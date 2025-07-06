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
)

type ServiceInventoryRepository interface {
	GetByID(context context.Context, id string, selectFields *string) (*models.Service, error)
	Create(context context.Context, serviceCreate *models.ServiceCreate) (*mongo.InsertOneResult, error)
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
