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
	"net/http"
)

type TmfServiceInventoryPkg struct {
	Server *restapi.Server
}

func NewTmfServiceInventory(l *core.Logger, db *core.DatabaseNeo4j, apiCore *core.API) *TmfServiceInventoryPkg {
	defer l.GetCore().Info("Initializing TMF SERVICE INVENTORY module")

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTmfServiceInventoryV42API(swaggerSpec)

	repo := service_inventory_repository.Neo4JServiceInventoryRepository{Db: db}
	serviceInventoryHandler := service_inventory_handler_v42.NewServiceInventoryHandler(&repo)

	// Register Handlers
	api.ServiceCreateServiceHandler = service.CreateServiceHandlerFunc(serviceInventoryHandler.CreateServiceHandler)
	api.ServiceRetrieveServiceHandler = service.RetrieveServiceHandlerFunc(serviceInventoryHandler.RetrieveServiceHandler)

	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.ConfigureAPI()

	// Register handlers in apiCore
	apiCore.GetRouter().Handle("/tmf-api/serviceInventory/v4/", http.StripPrefix("", server.GetHandler()))

	return &TmfServiceInventoryPkg{
		Server: server,
	}
}

func (tsi *TmfServiceInventoryPkg) GetTmfServiceInventory() *TmfServiceInventoryPkg {
	return tsi
}
