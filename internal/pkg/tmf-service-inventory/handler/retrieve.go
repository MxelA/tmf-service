package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
)

func (h *ServiceInventoryHandler) RetrieveServiceHandler(params service.RetrieveServiceParams) middleware.Responder {
	retrieveService, err := h.repo.GetByID(params.HTTPRequest.Context(), params.ID, params.Fields, params.GraphLookupDepth)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewRetrieveServiceInternalServerError().WithPayload(&errModel)
	}

	return service.NewRetrieveServiceOK().WithPayload(retrieveService)
}
