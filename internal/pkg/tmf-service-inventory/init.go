package tmf_service_inventory

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/handlers"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/loads"
	"log"
)

type TmfServiceInventory struct {
	API    *operations.TmfServiceInventoryV42API
	Server *restapi.Server
}

func NewTmfServiceInventory(l *core.Logger, db *core.DatabaseNeo4j) *TmfServiceInventory {
	defer l.GetCore().Info("Initializing TMF SERVICE INVENTORY")

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTmfServiceInventoryV42API(swaggerSpec)

	repo := service_inventory_repository.Neo4JServiceInventoryRepository{Db: db}
	serviceInventoryHandler := service_inventory_handler_v42.NewServiceInventoryHandler(&repo)

	api.ServiceCreateServiceHandler = service.CreateServiceHandlerFunc(serviceInventoryHandler.CreateServiceHandler)

	server := restapi.NewServer(api)

	return &TmfServiceInventory{
		API:    api,
		Server: server,
	}
}

func (tsi *TmfServiceInventory) GetTmfServiceInventory() *TmfServiceInventory {
	return tsi
}
