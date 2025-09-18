package tmf_service_order

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/middleware"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/event/publisher"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/handler"
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

func New(api *middleware.APIWrapper, db *core.DatabaseMongo, ps *core.PubSub, l *core.Logger, tr *core.Tracer) {

	// Initialize Mongo Repository
	mongoDb := db.GetCore()
	repo := &repository.MongoServiceOrderRepository{
		MongoCollection: mongoDb.Db.Collection(DbCollectionName),
		MongoClient:     mongoDb.Client,
		Logger:          l,
	}

	// Init PubSub
	pubSub := pub_sub.NewServiceOrderPubSub(ps, tr)
	pubSub.RegisterSubscribers()

	eventFactory := publisher.NewEventPublisher(pubSub, tr, l)
	// Initialize Handler
	serviceOrderHandler := handler.NewServiceOrderHandler(repo, l)
	serviceOrderOperators := registerOperators(serviceOrderHandler, eventFactory, pubSub, l)

	serviceOrderServer := restapi.NewServer(serviceOrderOperators)
	serviceOrderServer.ConfigureAPI()

	api.RegisterRoute("/tmf-api/serviceOrdering/v4/", http.StripPrefix("", serviceOrderServer.GetHandler()))
	api.RegisterRoute("/tmf-api/serviceOrdering/v4/docs", http.StripPrefix("/tmf-api/serviceOrdering/v4/docs", http.FileServer(http.Dir("internal/pkg/tmf-service-order/swagger/tmf641v4_2/doc"))))

	defer func(serviceOrderServer *restapi.Server) {
		_ = serviceOrderServer.Shutdown()
	}(serviceOrderServer)

	defer l.GetCore().Info("Initializing tmf-service-order package")

}

func registerOperators(serviceOrderHandler *handler.ServiceOrderHandler, eventPublisher *publisher.EventPublisher, pubSub *pub_sub.ServiceOrderPubSub, l *core.Logger) *operations.TmfServiceOrderV42API {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize Inventory Api
	serviceOrder := operations.NewTmfServiceOrderV42API(swaggerSpec)

	// Register create service order Handler with middleware
	createServiceOrderHandler := serviceOrderHandler.CreateServiceOrderHandler
	createServiceOrderHandler = local_middleware.SanitizeCreateServiceOrderDataMiddleware(createServiceOrderHandler)
	serviceOrder.ServiceOrderCreateServiceOrderHandler = service_order.CreateServiceOrderHandlerFunc(createServiceOrderHandler)

	serviceOrder.ServiceOrderDeleteServiceOrderHandler = service_order.DeleteServiceOrderHandlerFunc(serviceOrderHandler.DeleteServiceOrderHandler)
	serviceOrder.ServiceOrderListServiceOrderHandler = service_order.ListServiceOrderHandlerFunc(serviceOrderHandler.ListServiceOrderHandler)

	patchServiceOrderHandler := serviceOrderHandler.PatchServiceOrderHandler
	patchServiceOrderHandler = local_middleware.SendPatchServiceOrderEventMiddleware(eventPublisher, l, patchServiceOrderHandler)
	serviceOrder.ServiceOrderPatchServiceOrderHandler = service_order.PatchServiceOrderHandlerFunc(patchServiceOrderHandler)

	serviceOrder.ServiceOrderRetrieveServiceOrderHandler = service_order.RetrieveServiceOrderHandlerFunc(serviceOrderHandler.RetrieveServiceOrderHandler)

	return serviceOrder
}
