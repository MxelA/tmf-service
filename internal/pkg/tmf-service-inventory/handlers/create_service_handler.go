package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
)

func (h ServiceInventoryHandler) CreateServiceHandler(params service.CreateServiceParams) middleware.Responder {
	inventoryService, err := h.repo.Create(params.HTTPRequest.Context(), params.Service)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}

		return service.NewCreateServiceInternalServerError().WithPayload(&errModel)
	}

	return service.NewCreateServiceCreated().WithPayload(inventoryService)
}
