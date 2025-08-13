package middleware

import (
	"github.com/MxelA/tmf-service/internal/core"
	repository "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	validator "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/validators"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"strings"
	"time"
)

type CreateServiceInventoryFunc func(service.CreateServiceParams) middleware.Responder

func BusinessValidation(
	repo repository.ServiceInventoryRepository,
	next CreateServiceInventoryFunc,
) CreateServiceInventoryFunc {
	return func(params service.CreateServiceParams) middleware.Responder {
		// Call custom business validation
		errs := validator.ValidateBusinessRules(params, repo)
		if len(errs) > 0 {
			var messages []string
			for _, err := range errs {
				if err.Reason != nil {
					messages = append(messages, *err.Reason)
				}
			}
			combined := strings.Join(messages, "; ")
			code := "400"

			return service.NewCreateServiceBadRequest().WithPayload(&models.Error{
				Code:    &code,
				Reason:  &combined,
				Message: "Validation errors",
			})
		}

		return next(params)
	}
}

func BusinessValidationMiddleware(l *core.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.GetCore().Info("Started %s %s", r.Method, r.URL.Path)

			next.ServeHTTP(w, r)

			l.GetCore().Info("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
		})
	}
}
