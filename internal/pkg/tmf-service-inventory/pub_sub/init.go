package pub_sub

import (
	"github.com/MxelA/tmf-service/internal/core"
)

const (
	ServiceOrderStateChangeEventTopic string = "service_order_state_change"
)

type ServiceInventoryPubSub struct {
	PubSub *core.PubSub
	Logger *core.Logger
}

func NewServiceInventoryPubSub(ps *core.PubSub, l *core.Logger) *ServiceInventoryPubSub {
	return &ServiceInventoryPubSub{
		PubSub: ps,
		Logger: l,
	}
}

func (serviceOrderPubSub *ServiceInventoryPubSub) RegisterSubscribers() {
	serviceOrderPubSub.RegisterServiceOrderStateChangeSubscriber()
}
