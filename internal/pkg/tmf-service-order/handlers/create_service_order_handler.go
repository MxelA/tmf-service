package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
)

func (h *ServiceOrderHandler) CreateServiceOrderHandler(params service_order.CreateServiceOrderParams) middleware.Responder {
	serviceOrder, err := h.repo.Create(params.HTTPRequest.Context(), params.ServiceOrder)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}

		return service_order.NewCreateServiceOrderInternalServerError().WithPayload(&errModel)
	}

	return service_order.NewCreateServiceOrderCreated().WithPayload(serviceOrder)
}
