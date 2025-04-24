// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/events_subscription"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/notification_listeners_client_side"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
)

//go:generate swagger generate server --target ../../server --name TmfServiceInventoryV42 --spec ../../TMF638_Service_Inventory_Management_API_v4.2.0_beta_swagger.json --template-dir ./templates --principal interface{} --exclude-main

func configureFlags(api *operations.TmfServiceInventoryV42API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TmfServiceInventoryV42API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ServiceCreateServiceHandler == nil {
		api.ServiceCreateServiceHandler = service.CreateServiceHandlerFunc(func(params service.CreateServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.CreateService has not yet been implemented")
		})
	}
	if api.ServiceDeleteServiceHandler == nil {
		api.ServiceDeleteServiceHandler = service.DeleteServiceHandlerFunc(func(params service.DeleteServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.DeleteService has not yet been implemented")
		})
	}
	if api.ServiceListServiceHandler == nil {
		api.ServiceListServiceHandler = service.ListServiceHandlerFunc(func(params service.ListServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ListService has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceAttributeValueChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceAttributeValueChangeEventHandler = notification_listeners_client_side.ListenToServiceAttributeValueChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceAttributeValueChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceAttributeValueChangeEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceCreateEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceCreateEventHandler = notification_listeners_client_side.ListenToServiceCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceCreateEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceDeleteEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceDeleteEventHandler = notification_listeners_client_side.ListenToServiceDeleteEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceDeleteEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceDeleteEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceOperatingStatusChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceOperatingStatusChangeEventHandler = notification_listeners_client_side.ListenToServiceOperatingStatusChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceOperatingStatusChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceOperatingStatusChangeEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceStateChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceStateChangeEventHandler = notification_listeners_client_side.ListenToServiceStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceStateChangeEvent has not yet been implemented")
		})
	}
	if api.ServicePatchServiceHandler == nil {
		api.ServicePatchServiceHandler = service.PatchServiceHandlerFunc(func(params service.PatchServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.PatchService has not yet been implemented")
		})
	}
	if api.EventsSubscriptionRegisterListenerHandler == nil {
		api.EventsSubscriptionRegisterListenerHandler = events_subscription.RegisterListenerHandlerFunc(func(params events_subscription.RegisterListenerParams) middleware.Responder {
			return middleware.NotImplemented("operation events_subscription.RegisterListener has not yet been implemented")
		})
	}
	if api.ServiceRetrieveServiceHandler == nil {
		api.ServiceRetrieveServiceHandler = service.RetrieveServiceHandlerFunc(func(params service.RetrieveServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.RetrieveService has not yet been implemented")
		})
	}
	if api.EventsSubscriptionUnregisterListenerHandler == nil {
		api.EventsSubscriptionUnregisterListenerHandler = events_subscription.UnregisterListenerHandlerFunc(func(params events_subscription.UnregisterListenerParams) middleware.Responder {
			return middleware.NotImplemented("operation events_subscription.UnregisterListener has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
