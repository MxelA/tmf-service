package service_inventory_repository

import (
	"context"
	"github.com/MxelA/tmf-service/internal/core"
)

type ServiceInventoryRepository interface {
	GetByID(context context.Context, id string)
}

type Neo4JServiceInventoryRepository struct {
	Db *core.DatabaseNeo4j
}

func (repo *Neo4JServiceInventoryRepository) GetByID(context context.Context, id string) {

}
