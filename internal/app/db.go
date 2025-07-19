package app

import (
	"github.com/MxelA/tmf-service/internal/core"
	tmf_service_inventory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory"
	tmf_service_inventory_seeder "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/database/seeders"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/migrations"
)

func ServiceInventoryPkgSetMongoIndex() error {
	logger := core.NewLogger()
	db := core.NewDatabaseMongo(logger)

	return migrations.SetMongoIndex(db.GetCore().Db.Collection(tmf_service_inventory.CollectionName), logger)
}

func SeedServices(count int) error {
	logger := core.NewLogger()
	db := core.NewDatabaseMongo(logger)
	return tmf_service_inventory_seeder.SeedServices(db.GetCore().Db.Collection(tmf_service_inventory.CollectionName), count)
}

func SeedServicesWithRelationshipTo(count int) error {
	logger := core.NewLogger()
	db := core.NewDatabaseMongo(logger)
	return tmf_service_inventory_seeder.SeedServicesWithRelationshipTo(db.GetCore().Db.Collection(tmf_service_inventory.CollectionName), count)
}
