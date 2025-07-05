package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func (h ServiceInventoryHandler) RetrieveServiceHandler(params service.RetrieveServiceParams) middleware.Responder {
	retrieveService, err := h.repo.GetByID(params.HTTPRequest.Context(), params.ID, params.Fields)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service.NewRetrieveServiceInternalServerError().WithPayload(&errModel)
	}

	return service.NewRetrieveServiceOK().WithPayload(retrieveService)
}
