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
	tr     *core.Tracer
}

func NewServiceOrderPubSub(ps *core.PubSub, tr *core.Tracer) *ServiceOrderPubSub {
	return &ServiceOrderPubSub{
		pubSub: ps,
		tr:     tr,
	}
}

func (serviceOrderPubSub *ServiceOrderPubSub) RegisterSubscribers() {}
