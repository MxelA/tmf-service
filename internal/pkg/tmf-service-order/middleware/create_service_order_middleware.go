package middleware

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
)

type CreateServiceOrderFunc func(service_order.CreateServiceOrderParams) middleware.Responder

func SanitizeCreateServiceOrderDataMiddleware(
	l *core.Logger,
	next PatchServiceOrderFunc,
) PatchServiceOrderFunc {
	return func(params service_order.PatchServiceOrderParams) middleware.Responder {

		return next(params)
	}
}
