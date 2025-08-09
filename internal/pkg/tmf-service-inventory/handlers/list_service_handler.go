package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func (h *ServiceInventoryHandler) ListServiceHandler(params service.ListServiceParams) middleware.Responder {

	retrieveServiceOrders, retrieveServiceOrdersTotalCount, err := h.repo.GetAllPaginate(params.HTTPRequest.Context(), params.HTTPRequest, params.Fields, params.Offset, params.Limit)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service.NewListServiceInternalServerError().WithPayload(&errModel)
	}
	return service.NewListServiceOK().
		WithXTotalCount(*retrieveServiceOrdersTotalCount).
		WithXResultCount(int64(len(retrieveServiceOrders))).
		WithPayload(retrieveServiceOrders)
}
