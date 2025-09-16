package pub_sub

import (
	"context"
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	watermillmiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (serviceOrderPubSub *ServiceOrderPubSub) ServiceOrderStateChangePublisher(serviceOrderStateChangeEvent *models.ServiceOrderStateChangeEvent) {
	msg, _ := json.Marshal(&serviceOrderStateChangeEvent)
	eventMsg := message.NewMessage(watermill.NewUUID(), msg)
	watermillmiddleware.SetCorrelationID(*serviceOrderStateChangeEvent.CorrelationID, eventMsg)

	if err := serviceOrderPubSub.pubSub.GetCore().Publish(ServiceOrderStateChangeEventTopic, eventMsg); err != nil {
		panic(err)
	}
}

func (serviceOrderPubSub *ServiceOrderPubSub) ServiceOrderAttributeValueChangePublisher(serviceOrderStateChangeEvent *models.ServiceOrderAttributeValueChangeEvent, ctx context.Context) {
	msg, _ := json.Marshal(&serviceOrderStateChangeEvent)
	eventMsg := message.NewMessage(watermill.NewUUID(), msg)

	//watermillmiddleware.SetCorrelationID(*serviceOrderStateChangeEvent.CorrelationID, eventMsg)

	_, err := serviceOrderPubSub.tr.Trace(ctx, *serviceOrderStateChangeEvent.EventType, "service-order", func(ctx context.Context) error {
		span := trace.SpanFromContext(ctx)
		serviceOrderPubSub.tr.InjectTraceInEventMessage(ctx, eventMsg)

		span.AddEvent("Send Event:"+*serviceOrderStateChangeEvent.EventType, trace.WithAttributes(
			attribute.String("msg", string(msg)),
		))

		if err := serviceOrderPubSub.pubSub.GetCore().Publish(ServiceOrderAttributeValueChangeEventTopic, eventMsg); err != nil {
			span.AddEvent("Error:"+*serviceOrderStateChangeEvent.EventType, trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

}
