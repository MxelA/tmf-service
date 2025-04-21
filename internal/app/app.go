package app

import (
	"context"
	"fmt"
	"github.com/MxelA/tmf-service/internal/core"
	tmf_service_inventory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory"
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
	Addr                   string
	API                    *core.API
	DB                     *core.DatabaseNeo4j
	Logger                 *core.Logger
	TmfServiceInventoryPkg *tmf_service_inventory.TmfServiceInventoryPkg
	//PubSub *core.PubSub
}

func (app *app) Serve() {
	//api := app.API.GetCore()
	logger := app.Logger.GetCore()
	//pubSub := app.PubSub.GetCore()

	//mux := http.NewServeMux()

	//if app.TmfServiceInventoryPkg != nil {
	//	mux.Handle("/tmf-api/serviceInventory/v4/", http.StripPrefix("/tmf-api/serviceInventory/v4", app.TmfServiceInventoryPkg.Server.GetHandler()))
	//}

	mux := app.API.GetRouter()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	server := &http.Server{
		Addr:    app.Addr,
		Handler: app.API.GetRouter(),
	}

	go func() {
		logger.Info("App running", "address", app.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
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

	server.Shutdown(context.Background())
	//_ = api.ShutdownWithContext(context.Background())
	//_ = pubSub.Close()
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
