package app

import (
	"github.com/MxelA/tmf-service/internal/core"
	tmf_service_inventory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/migrations"
)

func ServiceInventoryPkgSetMongoIndex() error {
	logger := core.NewLogger()
	db := core.NewDatabaseMongo(logger)

	return migrations.SetMongoIndex(db.GetCore().Db.Collection(tmf_service_inventory.CollectionName), logger)
}
