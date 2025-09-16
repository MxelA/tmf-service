package middleware

import (
	"context"
	"github.com/MxelA/tmf-service/internal/core"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
)

// Middleware function type
type Middleware func(http.Handler) http.Handler

// ChainMiddleware primenjuje više middleware-a na handler
func ChainMiddleware(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

// APIWrapper save router and global middleware
type APIWrapper struct {
	Router    *http.ServeMux
	GlobalMws []Middleware
}

// NewAPIWrapper kreira novi wrapper
func NewAPIWrapper(router *http.ServeMux, globalMws ...Middleware) *APIWrapper {
	return &APIWrapper{
		Router:    router,
		GlobalMws: globalMws,
	}
}

// RegisterRoute register route with local and global middlewares
func (a *APIWrapper) RegisterRoute(pattern string, handler http.Handler, localMws ...Middleware) {
	// Combine global and local middlewares
	allMws := append(a.GlobalMws, localMws...)
	a.Router.Handle(pattern, ChainMiddleware(handler, allMws...))
}

// GetRouter
func (a *APIWrapper) GetRouter() http.Handler {
	return a.Router
}

func ApiLoggingMiddleware(l *core.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.GetCore().Info("Started", r.Method, r.URL.Path)

			next.ServeHTTP(w, r)

			l.GetCore().Info("Completed", r.Method, r.URL.Path, "in", time.Since(start))
		})
	}
}

func ApiTraceMiddleware(t *core.Tracer) func(http.Handler) http.Handler {
	type statusRecorder struct {
		http.ResponseWriter
		status int
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, _ = t.Trace(r.Context(), r.URL.Path, "api", func(ctx context.Context) error {
				span := trace.SpanFromContext(ctx)
				span.SetAttributes(
					attribute.String("http.method", r.Method),
					attribute.String("url.path", r.URL.Path),
				)

				span.AddEvent("Request received", trace.WithAttributes(
					attribute.String("payload", "{}"),
				))

				recorder := &statusRecorder{ResponseWriter: w, status: 200}

				next.ServeHTTP(w, r.WithContext(ctx))

				span.AddEvent("Response sent", trace.WithAttributes(
					attribute.Int("status", recorder.status),
				))

				return nil
			})
		})
	}
}
