package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func (h *ServiceOrderHandler) CreateServiceOrderHandler(params service_order.CreateServiceOrderParams) middleware.Responder {

	serviceOrder := models.ServiceOrder{}
	err := copier.Copy(&serviceOrder, params.ServiceOrder)
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

	uid := uuid.New().String()
	serviceOrder.ID = uid

	initState := models.ServiceOrderStateAcknowledged
	serviceOrder.State = &initState

	_, err = h.repo.Create(params.HTTPRequest.Context(), &serviceOrder)

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

	return service_order.NewCreateServiceOrderCreated().WithPayload(&serviceOrder)
}
