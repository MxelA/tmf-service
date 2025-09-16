package tmf_service_inventory

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/middleware"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/handlers"
	local_middleware "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/middleware"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/pub_sub"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/loads"
	"log"
	"net/http"
)

const DbCollectionName = "service_testing"

func New(api *middleware.APIWrapper, db *core.DatabaseMongo, pubSub *core.PubSub, l *core.Logger, tr *core.Tracer) {

	// Initialize Mongo Repository
	mongoDb := db.GetCore()
	repo := &repository.MongoServiceInventoryRepository{
		MongoCollection: mongoDb.Db.Collection(DbCollectionName),
		MongoClient:     mongoDb.Client,
		Logger:          l,
	}

	// Register Subscribers
	serviceInventoryPubSub := pub_sub.NewServiceInventoryPubSub(pubSub, repo, l, tr)
	serviceInventoryPubSub.RegisterSubscribers()

	// Initialize Handler
	serviceInventoryHandler := handler.NewServiceInventoryHandler(repo, l)
	serviceInventory := registerHandlers(serviceInventoryHandler, repo)

	inventoryServer := restapi.NewServer(serviceInventory)
	inventoryServer.ConfigureAPI()

	api.RegisterRoute("/tmf-api/serviceInventory/v4/", http.StripPrefix("", inventoryServer.GetHandler()))
	api.RegisterRoute("/tmf-api/serviceInventory/v4/docs", http.StripPrefix("/tmf-api/serviceInventory/v4/docs", http.FileServer(http.Dir("internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/doc"))))

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

	//Register Create Service Handler with middleware for business validation
	createServiceInventoryHandler := serviceInventoryHandler.CreateServiceHandler
	createServiceInventoryHandler = local_middleware.BusinessValidation(repo, createServiceInventoryHandler)
	inventory.ServiceCreateServiceHandler = service.CreateServiceHandlerFunc(createServiceInventoryHandler)

	// Register Handlers
	inventory.ServiceRetrieveServiceHandler = service.RetrieveServiceHandlerFunc(serviceInventoryHandler.RetrieveServiceHandler)
	inventory.ServiceListServiceHandler = service.ListServiceHandlerFunc(serviceInventoryHandler.ListServiceHandler)
	inventory.ServiceDeleteServiceHandler = service.DeleteServiceHandlerFunc(serviceInventoryHandler.DeleteServiceHandler)
	inventory.ServicePatchServiceHandler = service.PatchServiceHandlerFunc(serviceInventoryHandler.PatchServiceHandler)

	return inventory
}
