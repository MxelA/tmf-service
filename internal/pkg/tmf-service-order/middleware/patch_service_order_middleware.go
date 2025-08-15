package middleware

import (
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/pub_sub"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service/internal/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"time"
)

type PatchServiceOrderFunc func(service_order.PatchServiceOrderParams) middleware.Responder

func SendPatchServiceOrderEventMiddleware(
	serviceOrderPubSub *pub_sub.ServiceOrderPubSub,
	l *core.Logger,
	next PatchServiceOrderFunc,
) PatchServiceOrderFunc {
	return func(req service_order.PatchServiceOrderParams) middleware.Responder {
		resp := next(req)

		okResp, ok := resp.(*service_order.PatchServiceOrderOK)
		if !ok {
			return resp
		}

		patchMediaType := utils.DetectPatchMediaType(req.HTTPRequest.Header)
		switch *patchMediaType {
		case utils.JSONPatch:
			processJsonPatch(req, okResp, l)
		case utils.MergePatch:
			processMergePatch(serviceOrderPubSub, req, okResp, l)
		}

		l.GetCore().Info("SendPatchServiceOrderEventMiddleware")
		return resp
	}
}

func processJsonPatch(req service_order.PatchServiceOrderParams, okResp *service_order.PatchServiceOrderOK, l *core.Logger) {

	//marshal to json bytes
	raw, err := json.Marshal(req.ServiceOrder)
	if err != nil {
		return
	}

	var jpo []*models.JSONPatchOperation
	if err := json.Unmarshal(raw, &jpo); err != nil {
		return
	}

}

func processMergePatch(serviceOrderPubSub *pub_sub.ServiceOrderPubSub, req service_order.PatchServiceOrderParams, okResp *service_order.PatchServiceOrderOK, l *core.Logger) {
	raw, err := json.Marshal(req.ServiceOrder)
	if err != nil {
		return
	}

	var serviceOrderUpdate models.ServiceOrderUpdate
	if err = json.Unmarshal(raw, &serviceOrderUpdate); err != nil {
		return
	}

	if serviceOrderUpdate.State != okResp.Payload.State {
		sendServiceOrderStateChangeEvent(serviceOrderPubSub, okResp)
	} else {
		sendServiceOrderAttributeValueChangeEvent(okResp)
	}
}

func sendServiceOrderStateChangeEvent(serviceOrderPubSub *pub_sub.ServiceOrderPubSub, okResp *service_order.PatchServiceOrderOK) {
	id := uuid.New().String()
	eventType := "ServiceOrderStateChangeEvent"
	eventTime := strfmt.DateTime(time.Now().UTC())
	serviceOrderStateChangeEvent := models.ServiceOrderStateChangeEvent{
		CorrelationID: &id,
		EventType:     &eventType,
		Event: &models.ServiceOrderStateChangeEventPayload{
			ServiceOrder: okResp.Payload,
		},
		EventTime: &eventTime,
	}

	serviceOrderPubSub.ServiceOrderStateChangePublisher(&serviceOrderStateChangeEvent)
}

func sendServiceOrderAttributeValueChangeEvent(okResp *service_order.PatchServiceOrderOK) {
	id := uuid.New().String()
	serviceOrderStateEvent := models.ServiceOrderAttributeValueChangeEventPayload{
		ServiceOrder: okResp.Payload,
	}

	_ = models.ServiceOrderAttributeValueChangeEvent{
		CorrelationID: &id,
		Event:         &serviceOrderStateEvent,
	}
}
