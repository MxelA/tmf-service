package pub_sub

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/event/subscriber"
)

const (
	ServiceOrderAttributeValueChangeEventTopic string = "service_order_attribute_value_change"
)

type ServiceInventoryPubSub struct {
	pubSub          *core.PubSub
	eventSubscriber *subscriber.EventSubscriber
}

func NewServiceInventoryPubSub(ps *core.PubSub, eventSubscriber *subscriber.EventSubscriber) *ServiceInventoryPubSub {
	return &ServiceInventoryPubSub{
		pubSub:          ps,
		eventSubscriber: eventSubscriber,
	}
}

func (serviceOrderInventoryPubSub *ServiceInventoryPubSub) RegisterSubscribers() {
	serviceOrderInventoryPubSub.ServiceOrderAttributeValueChangeSubscriber()
}
