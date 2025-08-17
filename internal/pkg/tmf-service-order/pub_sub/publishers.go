package pub_sub

import (
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	watermillmiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

func (serviceOrderPubSub *ServiceOrderPubSub) ServiceOrderStateChangePublisher(serviceOrderStateChangeEvent *models.ServiceOrderStateChangeEvent) {
	msg, _ := json.Marshal(&serviceOrderStateChangeEvent)
	eventMsg := message.NewMessage(watermill.NewUUID(), msg)
	watermillmiddleware.SetCorrelationID(*serviceOrderStateChangeEvent.CorrelationID, eventMsg)

	if err := serviceOrderPubSub.pubSub.GetCore().Publish(ServiceOrderStateChangeEventTopic, eventMsg); err != nil {
		panic(err)
	}
}

func (serviceOrderPubSub *ServiceOrderPubSub) ServiceOrderAttributeValueChangePublisher(serviceOrderStateChangeEvent *models.ServiceOrderAttributeValueChangeEvent) {
	msg, _ := json.Marshal(&serviceOrderStateChangeEvent)
	eventMsg := message.NewMessage(watermill.NewUUID(), msg)
	watermillmiddleware.SetCorrelationID(*serviceOrderStateChangeEvent.CorrelationID, eventMsg)

	if err := serviceOrderPubSub.pubSub.GetCore().Publish(ServiceOrderAttributeValueChangeEventTopic, eventMsg); err != nil {
		panic(err)
	}
}
