package middleware

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
)

type CreateServiceOrderFunc func(service_order.CreateServiceOrderParams) middleware.Responder

func SanitizeCreateServiceOrderDataMiddleware(
	next CreateServiceOrderFunc,
) CreateServiceOrderFunc {
	return func(params service_order.CreateServiceOrderParams) middleware.Responder {

		for _, serviceOrderItem := range params.ServiceOrder.ServiceOrderItem {
			initState := models.ServiceOrderItemStateAcknowledged
			serviceOrderItem.State = &initState
		}

		return next(params)
	}
}
