package tmf_service_order

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/middleware"
	handler "github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/handlers"
	local_middleware "github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/middleware"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/pub_sub"
	repository "github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/repositories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/loads"
	"log"
	"net/http"
)

const DbCollectionName = "service_order"

func New(api *middleware.APIWrapper, db *core.DatabaseMongo, ps *core.PubSub, l *core.Logger) {

	// Initialize Mongo Repository
	mongoDb := db.GetCore()
	repo := &repository.MongoServiceOrderRepository{
		MongoCollection: mongoDb.Db.Collection(DbCollectionName),
		MongoClient:     mongoDb.Client,
		Logger:          l,
	}

	serviceOrderPubSub := pub_sub.NewServiceOrderPubSub(ps)
	serviceOrderPubSub.RegisterSubscribers()

	// Initialize Handler
	serviceOrderHandler := handler.NewServiceOrderHandler(repo, l)
	serviceOrderOperators := registerOperators(serviceOrderHandler, serviceOrderPubSub, l)

	serviceOrderServer := restapi.NewServer(serviceOrderOperators)
	serviceOrderServer.ConfigureAPI()

	api.RegisterRoute("/tmf-api/serviceOrdering/v4/", http.StripPrefix("", serviceOrderServer.GetHandler()))
	api.RegisterRoute("/tmf-api/serviceOrdering/v4/docs", http.StripPrefix("/tmf-api/serviceOrdering/v4/docs", http.FileServer(http.Dir("internal/pkg/tmf-service-order/swagger/tmf641v4_2/doc"))))

	defer func(serviceOrderServer *restapi.Server) {
		_ = serviceOrderServer.Shutdown()
	}(serviceOrderServer)

	defer l.GetCore().Info("Initializing tmf-service-order package")

}

func registerOperators(serviceOrderHandler *handler.ServiceOrderHandler, serviceOrderPubSub *pub_sub.ServiceOrderPubSub, l *core.Logger) *operations.TmfServiceOrderV42API {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize Inventory Api
	serviceOrder := operations.NewTmfServiceOrderV42API(swaggerSpec)

	// Register Handlers
	serviceOrder.ServiceOrderCreateServiceOrderHandler = service_order.CreateServiceOrderHandlerFunc(serviceOrderHandler.CreateServiceOrderHandler)
	serviceOrder.ServiceOrderDeleteServiceOrderHandler = service_order.DeleteServiceOrderHandlerFunc(serviceOrderHandler.DeleteServiceOrderHandler)
	serviceOrder.ServiceOrderListServiceOrderHandler = service_order.ListServiceOrderHandlerFunc(serviceOrderHandler.ListServiceOrderHandler)

	patchServiceOrderHandler := serviceOrderHandler.PatchServiceOrderHandler
	patchServiceOrderHandler = local_middleware.SendPatchServiceOrderEventMiddleware(serviceOrderPubSub, l, patchServiceOrderHandler)
	serviceOrder.ServiceOrderPatchServiceOrderHandler = service_order.PatchServiceOrderHandlerFunc(patchServiceOrderHandler)

	serviceOrder.ServiceOrderRetrieveServiceOrderHandler = service_order.RetrieveServiceOrderHandlerFunc(serviceOrderHandler.RetrieveServiceOrderHandler)

	return serviceOrder
}
