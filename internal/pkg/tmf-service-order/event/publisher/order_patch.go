package publisher

import (
	"context"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"time"
)

func (ef *EventPublisher) SendServiceOrderStateChangeEvent(serviceOrder *models.ServiceOrder, ctx context.Context) {
	id := uuid.New().String()
	eventType := "ServiceOrderStateChangeEvent"
	eventTime := strfmt.DateTime(time.Now().UTC())

	serviceOrderStateChangeEvent := models.ServiceOrderStateChangeEvent{
		CorrelationID: &id,
		EventType:     &eventType,
		Event: &models.ServiceOrderStateChangeEventPayload{
			ServiceOrder: serviceOrder,
		},
		EventTime: &eventTime,
	}

	ef.pubSub.ServiceOrderStateChangePublisher(&serviceOrderStateChangeEvent)
}

func (ef *EventPublisher) SendServiceOrderAttributeValueChangeEvent(serviceOrder *models.ServiceOrder, ctx context.Context) {
	id := uuid.New().String()
	eventType := "ServiceOrderAttributeValueChangeEvent"
	eventTime := strfmt.DateTime(time.Now().UTC())

	serviceOrderAttributeValueChange := models.ServiceOrderAttributeValueChangeEvent{
		CorrelationID: &id,
		EventType:     &eventType,
		EventTime:     &eventTime,
		Event: &models.ServiceOrderAttributeValueChangeEventPayload{
			ServiceOrder: serviceOrder,
		},
	}

	ef.pubSub.ServiceOrderAttributeValueChangePublisher(&serviceOrderAttributeValueChange, ctx)
}
