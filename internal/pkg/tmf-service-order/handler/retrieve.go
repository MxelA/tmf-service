package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
)

func (h *ServiceOrderHandler) RetrieveServiceOrderHandler(params service_order.RetrieveServiceOrderParams) middleware.Responder {
	retrieveServiceOrder, err := h.repo.GetByID(params.HTTPRequest.Context(), params.ID, params.Fields)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewRetrieveServiceOrderInternalServerError().WithPayload(&errModel)
	}

	return service_order.NewRetrieveServiceOrderOK().WithPayload(retrieveServiceOrder)
}
