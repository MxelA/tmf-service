package migrations

import (
	"context"
	"github.com/MxelA/tmf-service/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetMongoIndex(collection *mongo.Collection, logger *core.Logger) error {
	//indexModel := mongo.IndexModel{
	//	Keys: bson.D{
	//		{"serviceRelationship.service.id", 1},
	//	},
	//	// Add index name:
	//	Options: options.Index().SetName("idx_serviceRelationship_service_id"),
	//}

	indexModels := []mongo.IndexModel{
		{
			Keys: bson.D{
				{"serviceRelationship.service.id", 1},
			},
			// Add index name:
			//Options: options.Index().SetName("idx_serviceRelationship_service_id"),
		},
		{
			Keys: bson.D{
				{"category", 1},
			},
			// Add index name:
			//Options: options.Index().SetName("idx_category"),
		},
	}
	name, err := collection.Indexes().CreateMany(context.TODO(), indexModels)
	if err != nil {
		return err
	}

	logger.GetCore().Info("✅ Index created:", name)
	return nil
}
