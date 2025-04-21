package core

import (
	"net/http"
)

// API responsible to deliver communication of
// application and client via http transport.
type API struct {
	router *http.ServeMux
}

func NewApi(l *Logger) *API {
	defer l.GetCore().Info("API Initialized")

	return &API{router: http.NewServeMux()}
}

func (r *API) GetRouter() *http.ServeMux {
	return r.router
}
