package middleware

import (
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/event/publisher"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service/internal/utils"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/go-openapi/runtime/middleware"
)

type PatchServiceOrderFunc func(service_order.PatchServiceOrderParams) middleware.Responder

func SendPatchServiceOrderEventMiddleware(
	eventPublisher *publisher.EventPublisher,
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
			processJsonPatch(eventPublisher, req, okResp)
		case utils.MergePatch:
			processMergePatch(eventPublisher, req, okResp)
		}

		l.GetCore().Info("SendPatchServiceOrderEventMiddleware")
		return resp
	}
}

func processJsonPatch(eventPublisher *publisher.EventPublisher, req service_order.PatchServiceOrderParams, okResp *service_order.PatchServiceOrderOK) {

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
		eventPublisher.SendServiceOrderStateChangeEvent(okResp.Payload, req.HTTPRequest.Context())
	}

	if sendOrderAttributeValueChangeEvent {
		eventPublisher.SendServiceOrderAttributeValueChangeEvent(okResp.Payload, req.HTTPRequest.Context())
	}
}

func processMergePatch(eventPublisher *publisher.EventPublisher, req service_order.PatchServiceOrderParams, okResp *service_order.PatchServiceOrderOK) {
	raw, err := json.Marshal(req.ServiceOrder)
	if err != nil {
		return
	}

	var serviceOrderUpdate models.ServiceOrderUpdate
	if err = json.Unmarshal(raw, &serviceOrderUpdate); err != nil {
		return
	}

	if utils.IsOnlyFieldSet(serviceOrderUpdate, "State") {
		eventPublisher.SendServiceOrderStateChangeEvent(okResp.Payload, req.HTTPRequest.Context())
	} else {
		if serviceOrderUpdate.State != nil {
			eventPublisher.SendServiceOrderStateChangeEvent(okResp.Payload, req.HTTPRequest.Context())
		}
		eventPublisher.SendServiceOrderAttributeValueChangeEvent(okResp.Payload, req.HTTPRequest.Context())
	}
}
