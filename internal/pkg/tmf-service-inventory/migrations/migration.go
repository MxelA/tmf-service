package migrations

import (
	"context"
	"github.com/MxelA/tmf-service/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetMongoIndex(collection *mongo.Collection, logger *core.Logger) error {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"serviceRelationship.service.id", 1},
		},
		// Add index name:
		Options: options.Index().SetName("idx_serviceRelationship_service_id"),
	}

	name, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}

	logger.GetCore().Info("✅ Index created:", name)
	return nil
}
