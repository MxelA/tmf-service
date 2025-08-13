package app

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/middleware"
	tmf_service_inventory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory"
	tmf_service_order "github.com/MxelA/tmf-service/internal/pkg/tmf-service-order"
)

// initCore initialize core packages of application
// such as router, logger, database, etc.
func (app *app) initCore() {
	logger := core.NewLogger()
	app.Logger = logger

	db := core.NewDatabaseMongo(logger)
	app.DBMongo = db

	api := core.NewApi(logger)
	apiWrapper := middleware.NewAPIWrapper(api.GetRouter(),
		middleware.ApiLoggingMiddleware(logger),
	)
	app.API = api
	app.APIWrapper = apiWrapper

	pubSub := core.NewPubSub(logger)
	middleware.InitPubSubMiddleware(pubSub)
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
		//pubSub = app.PubSub
		logger = app.Logger
	)

	tmf_service_inventory.New(apiWrapper, db, logger)
	tmf_service_order.New(apiWrapper, db, logger)

	defer logger.GetCore().Info("Initialize packages done!")
}
