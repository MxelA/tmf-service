package app

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/middleware"
	tmf_service_inventory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory"
	tmf_service_order "github.com/MxelA/tmf-service/internal/pkg/tmf-service-order"
	watermill_middleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

// initCore initialize core packages of application
// such as router, logger, database, etc.
func (app *app) initCore() {

	// Initialize Logger
	logger := core.NewLogger()
	app.Logger = logger

	// Initialize Tracer
	tracer := core.NewJaegerTracer(logger)
	app.Tracer = tracer

	// Initialize Mongo DB
	db := core.NewDatabaseMongo(logger)
	app.DBMongo = db

	// Initialize Web Api
	api := core.NewApi(logger)
	apiWrapper := middleware.NewAPIWrapper(api.GetRouter(),
		middleware.ApiTraceMiddleware(tracer),
	)
	app.API = api
	app.APIWrapper = apiWrapper

	// Initialize PubSub
	pubSub := core.NewPubSub(logger)
	middleware.InitPubSubMiddleware(pubSub,
		watermill_middleware.CorrelationID,
		watermill_middleware.Recoverer,
	)
	app.PubSub = pubSub

	defer logger.GetCore().Info("Initialize dependencies done!")
}

// initPackages initialize all the packages inside the pkg directory.
// This function act as single source where all
// packages should be initialized.
func (app *app) initPackages() {
	var (
		db = app.DBMongo
		//api        = app.API
		apiWrapper = app.APIWrapper
		pubSub     = app.PubSub
		logger     = app.Logger
		tracer     = app.Tracer
	)

	tmf_service_inventory.New(apiWrapper, db, pubSub, logger, tracer)
	tmf_service_order.New(apiWrapper, db, pubSub, logger, tracer)

	defer logger.GetCore().Info("Initialize packages done!")
}
