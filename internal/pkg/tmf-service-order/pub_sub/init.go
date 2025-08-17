package pub_sub

import (
	"github.com/MxelA/tmf-service/internal/core"
)

const (
	ServiceOrderStateChangeEventTopic          string = "service_order_state_change"
	ServiceOrderAttributeValueChangeEventTopic string = "service_order_attribute_value_change"
)

type ServiceOrderPubSub struct {
	pubSub *core.PubSub
}

func NewServiceOrderPubSub(ps *core.PubSub) *ServiceOrderPubSub {
	return &ServiceOrderPubSub{
		pubSub: ps,
	}
}

func (serviceOrderPubSub *ServiceOrderPubSub) RegisterSubscribers() {}
