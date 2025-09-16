package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
)

func (h *ServiceOrderHandler) DeleteServiceOrderHandler(params service_order.DeleteServiceOrderParams) middleware.Responder {
	_, err := h.repo.Delete(params.HTTPRequest.Context(), params.ID)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}

		return service_order.NewDeleteServiceOrderInternalServerError().WithPayload(&errModel)
	}

	return service_order.NewDeleteServiceOrderNoContent()
}
