package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	//"syscall"
)

type App interface {
	Serve()
}

// app contain all the core dependency and should
// be implement at application startup.
type app struct {
	Addr       string
	API        *core.API
	APIWrapper *middleware.APIWrapper
	DBMongo    *core.DatabaseMongo
	Logger     *core.Logger
	PubSub     *core.PubSub
}

// // New would implement the App interface by
// // initialize the app's core dependencies and packages.
func New(host, port string) App {
	var app = new(app)

	app.Addr = fmt.Sprintf("%s:%s", host, port)
	app.initCore()
	app.initPackages()

	return app
}

func (app *app) Serve() {
	api := app.API.GetCore()
	logger := app.Logger.GetCore()
	pubSub := app.PubSub.GetCore()

	//server := &http.Server{
	//	Addr:    app.Addr,
	//	Handler: app.API.GetRouter(),
	//}

	api.Addr = app.Addr
	api.Handler = app.API.GetRouter()

	go func() {
		logger.Info("App running", "address", app.Addr)
		if err := api.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			// Unexpected error, log and exit
			fmt.Printf("ListenAndServe error: %v\n", err)
			os.Exit(1)
		}
	}()

	var sig os.Signal
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	sig = <-c                                       // This blocks the main thread until an interrupt is received

	logger.Info("Signal received", "signal", sig.String())
	logger.Info("Shutting down app, waiting background process to finish")

	defer logger.Info("App shutdown")

	_ = api.Shutdown(context.Background())
	_ = pubSub.Close()
	//_ = api.ShutdownWithContext(context.Background())

}
