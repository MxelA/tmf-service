package tmf_service_inventory

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/handlers"
	middleware "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/middlewares"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/loads"
	"log"
	"net/http"
)

const DbCollectionName = "serviceInventory"

func New(api *core.API, db *core.DatabaseMongo, l *core.Logger) {

	// Initialize Mongo Repository
	repo := &repository.MongoServiceInventoryRepository{
		MongoCollection: db.GetCore().Db.Collection(DbCollectionName),
		Logger:          l,
	}
	// Initialize Handler
	serviceInventoryHandler := handler.NewServiceInventoryHandler(repo, l)

	serviceInventory := registerHandlers(serviceInventoryHandler, repo)

	inventoryServer := restapi.NewServer(serviceInventory)
	inventoryServer.ConfigureAPI()

	api.GetRouter().Handle("/tmf-api/serviceInventory/v4/", http.StripPrefix("", inventoryServer.GetHandler()))
	api.GetRouter().Handle("/tmf-api/serviceInventory/v4/docs", http.StripPrefix("/tmf-api/serviceInventory/v4/docs", http.FileServer(http.Dir("internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/doc"))))

	defer func(inventoryServer *restapi.Server) {
		_ = inventoryServer.Shutdown()
	}(inventoryServer)

	defer l.GetCore().Info("Initializing tmf-service-inventory package")

}

func registerHandlers(serviceInventoryHandler *handler.ServiceInventoryHandler, repo *repository.MongoServiceInventoryRepository) *operations.TmfServiceInventoryV42API {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize Inventory Api
	inventory := operations.NewTmfServiceInventoryV42API(swaggerSpec)

	//Register Create Service Handler
	createServiceInventoryHandler := serviceInventoryHandler.CreateServiceHandler
	createServiceInventoryHandler = middleware.BusinessValidation(repo, createServiceInventoryHandler)
	inventory.ServiceCreateServiceHandler = service.CreateServiceHandlerFunc(createServiceInventoryHandler)

	// Register Handlers
	inventory.ServiceRetrieveServiceHandler = service.RetrieveServiceHandlerFunc(serviceInventoryHandler.RetrieveServiceHandler)
	inventory.ServiceListServiceHandler = service.ListServiceHandlerFunc(serviceInventoryHandler.ListServiceHandler)
	inventory.ServiceDeleteServiceHandler = service.DeleteServiceHandlerFunc(serviceInventoryHandler.DeleteServiceHandler)

	return inventory
}
