package app

import (
	"github.com/MxelA/tmf-service/internal/core"
	tmf_service_inventory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/database/migrations"
	tmf_service_inventory_seeder "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/database/seeders"
)

func ServiceInventoryPkgSetMongoIndex() error {
	logger := core.NewLogger()
	db := core.NewDatabaseMongo(logger)

	return migrations.CreateMongoIndex(db.GetCore().Db.Collection(tmf_service_inventory.DbCollectionName), logger)
}

func SeedServices(count int) error {
	logger := core.NewLogger()
	db := core.NewDatabaseMongo(logger)
	return tmf_service_inventory_seeder.SeedServices(db.GetCore().Db.Collection(tmf_service_inventory.DbCollectionName), count)
}

func SeedServicesWithRelationshipTo(count int) error {
	logger := core.NewLogger()
	db := core.NewDatabaseMongo(logger)
	return tmf_service_inventory_seeder.SeedServicesWithRelationshipTo(db.GetCore().Db.Collection(tmf_service_inventory.DbCollectionName), count)
}
