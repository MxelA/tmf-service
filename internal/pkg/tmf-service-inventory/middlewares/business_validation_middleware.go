package middleware

import (
	repository "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	validator "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/validators"
	"github.com/go-openapi/runtime/middleware"
	"strings"
)

type CreateServiceInventoryFunc func(service.CreateServiceParams) middleware.Responder

func WithBusinessValidation(
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
