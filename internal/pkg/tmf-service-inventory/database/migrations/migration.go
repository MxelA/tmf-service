package migrations

import (
	"context"
	"github.com/MxelA/tmf-service/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoIndex(collection *mongo.Collection, logger *core.Logger) error {
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.D{
				{"id", 1},
			},
			Options: options.Index().SetName("idx_id"), // Add index name:
		},
		{
			Keys: bson.D{
				{"category", 1},
			},
			Options: options.Index().SetName("idx_category"), // Add index name:
		},
		{
			Keys: bson.D{
				{"serviceRelationship.service.id", 1},
			},
			Options: options.Index().SetName("idx_serviceRelationship_service_id"), // Add index name:
		},
	}

	name, err := collection.Indexes().CreateMany(context.TODO(), indexModels)
	if err != nil {
		return err
	}

	logger.GetCore().Info("✅ Index created:", name)
	return nil
}
