package subscriber

import (
	"context"
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (eventSubscriber *EventSubscriber) ServiceOrderChangeSubscriberHandler(msg *message.Message) error {
	
	traceContext := eventSubscriber.tracer.ExtractTraceFromEventMessage(msg)

	_, err := eventSubscriber.tracer.Trace(traceContext, "ServiceOrderAttributeValueChangeEventTopic", "service-inventory", func(ctx context.Context) error {
		span := trace.SpanFromContext(ctx)
		span.AddEvent("Receive message from ServiceOrderAttributeValueChangeEvent")
		services, err := parseServicesFromServiceOrderEventMsg(msg)
		if err != nil {
			span.AddEvent("Error parseServiceOrderStateChangeMsg", trace.WithAttributes(
				attribute.String("error", err.Error()),
			))
			return err
		}
		span.AddEvent("Finish parsing ServiceOrderAttributeValueChangeEvent")
		for _, service := range services {
			if service.ID != nil {
				serviceInventory, _ := eventSubscriber.repo.GetByID(traceContext, *service.ID, nil, nil)
				for _, serviceOrderItem := range serviceInventory.ServiceOrderItem {
					if serviceOrderItem != nil && *serviceOrderItem.ServiceOrderID == *service.ServiceOrderItem[0].ServiceOrderID && serviceOrderItem.ItemID == service.ServiceOrderItem[0].ItemID {
						return nil
					}
				}
				service.ServiceOrderItem = append(serviceInventory.ServiceOrderItem, service.ServiceOrderItem[0])

				_, err := eventSubscriber.repo.Update(traceContext, *service.ID, service)
				if err != nil {
					return err
				}
			} else {
				serviceCreate := models.ServiceCreate{}
				_ = copier.Copy(&serviceCreate, &service)
				_, _ = eventSubscriber.repo.Create(traceContext, &serviceCreate)
			}
		}

		return nil
	})

	return err
}

func parseServicesFromServiceOrderEventMsg(msg *message.Message) ([]*models.Service, error) {
	var services []*models.Service
	eventWrapper := EventWrapper{}

	if err := json.Unmarshal(msg.Payload, &eventWrapper); err != nil {
		return services, err
	}

	for _, orderItem := range eventWrapper.Event.ServiceOrder.ServiceOrderItem {
		if orderItem.State == "completed" {
			service := orderItem.Service
			service.ServiceOrderItem = []*models.RelatedServiceOrderItem{
				{
					ServiceOrderID:   &eventWrapper.Event.ServiceOrder.ID,
					ItemID:           orderItem.ID,
					ServiceOrderHref: &eventWrapper.Event.ServiceOrder.Href,
					ItemAction:       orderItem.Action,
				},
			}
			if orderItem.Action == models.OrderItemActionTypeAdd && service.State == "" {
				service.State = models.ServiceStateTypeInactive
			}
			services = append(services, &service)
		}
	}

	return services, nil
}
