package middleware

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

type PatchServiceOrderFunc func(service_order.PatchServiceOrderParams) middleware.Responder

func SendPatchServiceOrderEventMiddleware(
	l *core.Logger,
	next PatchServiceOrderFunc,
) PatchServiceOrderFunc {
	return func(params service_order.PatchServiceOrderParams) middleware.Responder {
		resp := next(params)

		okResp, ok := resp.(*service_order.PatchServiceOrderOK)
		if !ok {
			return resp
		}

		id := uuid.New().String()
		serviceOrderStateEvent := models.ServiceOrderStateChangeEventPayload{
			ServiceOrder: okResp.Payload,
		}

		_ = models.ServiceOrderStateChangeEvent{
			CorrelationID: &id,
			Event:         &serviceOrderStateEvent,
		}

		l.GetCore().Info("Sending patch service order event after next call")
		return resp
	}
}
