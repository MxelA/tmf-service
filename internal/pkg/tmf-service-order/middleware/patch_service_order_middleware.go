package middleware

import (
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/pub_sub"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service/internal/utils"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
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
			processJsonPatch(serviceOrderPubSub, req, okResp, l)
		case utils.MergePatch:
			processMergePatch(serviceOrderPubSub, req, okResp, l)
		}

		l.GetCore().Info("SendPatchServiceOrderEventMiddleware")
		return resp
	}
}

func processJsonPatch(serviceOrderPubSub *pub_sub.ServiceOrderPubSub, req service_order.PatchServiceOrderParams, okResp *service_order.PatchServiceOrderOK, l *core.Logger) {

	//marshal to json bytes
	raw, err := json.Marshal(req.ServiceOrder)
	if err != nil {
		return
	}

	var jpo []*models.JSONPatchOperation
	if err := json.Unmarshal(raw, &jpo); err != nil {
		return
	}

	sendOrderStateChangeEvent := false
	sendOrderAttributeValueChangeEvent := false

	patchOperations, _ := jsonpatch.DecodePatch(raw)
	for _, p := range patchOperations {
		path, _ := p.From()

		if path == "state" {
			sendOrderStateChangeEvent = true
		} else {
			sendOrderAttributeValueChangeEvent = true
		}
	}

	if sendOrderStateChangeEvent {
		sendServiceOrderStateChangeEvent(serviceOrderPubSub, req, okResp)
	}

	if sendOrderAttributeValueChangeEvent {
		sendServiceOrderAttributeValueChangeEvent(serviceOrderPubSub, req, okResp)
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

	if utils.IsOnlyFieldSet(serviceOrderUpdate, "State") {
		sendServiceOrderStateChangeEvent(serviceOrderPubSub, req, okResp)
	} else {
		if serviceOrderUpdate.State != nil {
			sendServiceOrderStateChangeEvent(serviceOrderPubSub, req, okResp)
		}
		sendServiceOrderAttributeValueChangeEvent(serviceOrderPubSub, req, okResp)
	}

}

func sendServiceOrderStateChangeEvent(serviceOrderPubSub *pub_sub.ServiceOrderPubSub, req service_order.PatchServiceOrderParams, okResp *service_order.PatchServiceOrderOK) {
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

	tracer := otel.Tracer("serviceOrdering")
	_, span := tracer.Start(req.HTTPRequest.Context(), eventType)
	defer span.End()

	serviceOrderPubSub.ServiceOrderStateChangePublisher(&serviceOrderStateChangeEvent)
}

func sendServiceOrderAttributeValueChangeEvent(serviceOrderPubSub *pub_sub.ServiceOrderPubSub, req service_order.PatchServiceOrderParams, okResp *service_order.PatchServiceOrderOK) {
	id := uuid.New().String()
	eventType := "ServiceOrderAttributeValueChangeEvent"
	eventTime := strfmt.DateTime(time.Now().UTC())

	serviceOrderAttributeValueChange := models.ServiceOrderAttributeValueChangeEvent{
		CorrelationID: &id,
		EventType:     &eventType,
		EventTime:     &eventTime,
		Event: &models.ServiceOrderAttributeValueChangeEventPayload{
			ServiceOrder: okResp.Payload,
		},
	}

	serviceOrderPubSub.ServiceOrderAttributeValueChangePublisher(&serviceOrderAttributeValueChange, req.HTTPRequest.Context())
}
