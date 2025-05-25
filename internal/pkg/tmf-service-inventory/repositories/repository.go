package repository

import (
	"context"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_0/server/models"
	"github.com/MxelA/tmf-service/internal/utils"
)

type ServiceInventoryRepository interface {
	GetByID(context context.Context, id string)
	Create(context context.Context, serviceCreate *models.ServiceCreate) error
}

type Neo4JServiceInventoryRepository struct {
	Db     *core.DatabaseNeo4j
	Logger *core.Logger
}

func (repo *Neo4JServiceInventoryRepository) GetByID(context context.Context, id string) {

}
func (repo *Neo4JServiceInventoryRepository) Create(context context.Context, serviceCreate *models.ServiceCreate) error {
	a, err := utils.ToSlice(serviceCreate)

	utils.PrettyPrint(a)

	if err != nil {
		return err
	}

	repo.Db.GetCore()

	return nil
}
