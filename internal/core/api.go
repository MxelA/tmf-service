package core

import (
	"net/http"
)

// API responsible to deliver communication of
// application and client via http transport.
type API struct {
	core *http.Server
	mux  *http.ServeMux
}

func NewApi(l *Logger) *API {
	defer l.GetCore().Info("API Initialized")

	return &API{
		core: &http.Server{},
		mux:  http.NewServeMux(),
	}
}

func (r *API) GetRouter() *http.ServeMux {
	return r.mux
}

func (r *API) GetCore() *http.Server {
	return r.core
}
