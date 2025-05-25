package tmf_service_inventory

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/handlers"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_0/server/restapi"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_0/server/restapi/operations"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_0/server/restapi/operations/service"
	"github.com/go-openapi/loads"
	"log"
)

type TmfServiceInventoryPkg struct {
	Server *restapi.Server
}

func NewTmfServiceInventory(l *core.Logger, db *core.DatabaseNeo4j) *TmfServiceInventoryPkg {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTmfServiceInventoryV40API(swaggerSpec)

	repo := repository.Neo4JServiceInventoryRepository{Db: db, Logger: l}
	serviceInventoryHandler := handler.NewServiceInventoryHandler(&repo)

	// Register Handlers
	api.ServiceCreateServiceHandler = service.CreateServiceHandlerFunc(serviceInventoryHandler.CreateServiceHandler)
	api.ServiceRetrieveServiceHandler = service.RetrieveServiceHandlerFunc(serviceInventoryHandler.RetrieveServiceHandler)

	server := restapi.NewServer(api)
	server.ConfigureAPI()

	defer server.Shutdown()
	defer l.GetCore().Info("Initializing tmf-service-inventory package")

	return &TmfServiceInventoryPkg{
		Server: server,
	}
}

func (tsi *TmfServiceInventoryPkg) GetTmfServiceInventory() *TmfServiceInventoryPkg {
	return tsi
}
