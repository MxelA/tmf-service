package app

import (
	"github.com/MxelA/tmf-service/internal/core"
	tmf_service_inventory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory"
	"net/http"
)

// initCore initialize core packages of application
// such as router, logger, database, etc.
func (app *app) initCore() {
	logger := core.NewLogger()
	app.Logger = logger

	db := core.NewDatabaseNeo4j(logger)
	app.DB = db

	api := core.NewApi(logger)
	app.API = api

	//pubSub := core.NewPubSub(logger)
	//middleware.InitPubSubMiddleware(pubSub)
	//app.PubSub = pubSub

	defer logger.GetCore().Info("Initialize dependencies done!")
}

// initPackages initialize all the packages inside the pkg directory.
// This function act as single source where all
// packages should be initialized.
func (app *app) initPackages() {
	var (
		db = app.DB
		//api = app.API
		//pubSub = app.PubSub
		logger = app.Logger
	)

	// initialize tmf-service-inventory package
	tmfServiceInventoryPkg := tmf_service_inventory.NewTmfServiceInventory(logger, db)
	app.API.GetRouter().Handle("/tmf-api/serviceInventory/v4/", http.StripPrefix("", tmfServiceInventoryPkg.Server.GetHandler()))

	defer logger.GetCore().Info("Initialize packages done!")
}
