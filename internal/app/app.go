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
	"time"
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
	Tracer     *core.Tracer
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
	pubSubRouter := app.PubSub.GetRouter()

	api.Addr = app.Addr
	api.Handler = app.API.GetRouter()

	// Start HTTP server
	httpErrChan := make(chan error, 1)
	go func() {
		logger.Info("App running", "address", app.Addr)
		if err := api.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			httpErrChan <- fmt.Errorf("ListenAndServe error: %w", err)
		} else {
			httpErrChan <- nil
		}
	}()

	// Start Pub/Sub router
	pubSubErrChan := make(chan error, 1)
	go func() {
		if err := pubSubRouter.Run(context.Background()); err != nil {
			pubSubErrChan <- fmt.Errorf("PubSub Router error: %w", err)
		} else {
			pubSubErrChan <- nil
		}
	}()

	// Wait signal for shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan

	logger.Info("Signal received", "signal", sig.String())
	logger.Info("Shutting down app, waiting for background processes")
	defer logger.Info("App shutdown complete")

	// Shutdown HTTP server with timeout
	httpCtx, httpCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer httpCancel()

	if err := api.Shutdown(httpCtx); err != nil {
		logger.Error("HTTP server shutdown error", "err", err)
	}
	logger.Info("HTTP server shutdown complete")

	// Shutdown Pub/Sub with timeout (5s)
	pubSubDone := make(chan struct{})
	go func() {
		_ = pubSub.Close()
		close(pubSubDone)
	}()

	select {
	case <-pubSubDone:
		logger.Info("Pub/Sub closed successfully")
	case <-time.After(5 * time.Second):
		logger.Warn("Pub/Sub shutdown timed out")
	}

}
