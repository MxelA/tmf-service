package middleware

import (
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/pub_sub"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service/internal/utils"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	watermillmiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"time"
)

type PatchServiceOrderFunc func(service_order.PatchServiceOrderParams) middleware.Responder

func SendPatchServiceOrderEventMiddleware(
	ps *core.PubSub,
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
			processMergePatch(ps, req, okResp, l)
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

func processMergePatch(ps *core.PubSub, req service_order.PatchServiceOrderParams, okResp *service_order.PatchServiceOrderOK, l *core.Logger) {
	raw, err := json.Marshal(req.ServiceOrder)
	if err != nil {
		return
	}

	var serviceOrderUpdate models.ServiceOrderUpdate
	if err = json.Unmarshal(raw, &serviceOrderUpdate); err != nil {
		return
	}

	if serviceOrderUpdate.State != okResp.Payload.State {
		sendServiceOrderStateChangeEvent(ps, okResp)
	} else {
		sendServiceOrderAttributeValueChangeEvent(okResp)
	}
}

func sendServiceOrderStateChangeEvent(ps *core.PubSub, okResp *service_order.PatchServiceOrderOK) {
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

	msg, _ := json.Marshal(&serviceOrderStateChangeEvent)
	eventMsg := message.NewMessage(watermill.NewUUID(), msg)
	watermillmiddleware.SetCorrelationID(*serviceOrderStateChangeEvent.CorrelationID, eventMsg)

	if err := ps.GetCore().Publish(pub_sub.ServiceOrderStateChangeEventTopic, eventMsg); err != nil {
		panic(err)
	}
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
