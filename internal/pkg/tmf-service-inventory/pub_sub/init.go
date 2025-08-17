package pub_sub

import (
	"github.com/MxelA/tmf-service/internal/core"
	repository "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
)

const (
	ServiceOrderStateChangeEventTopic string = "service_order_state_change"
)

type ServiceInventoryPubSub struct {
	PubSub *core.PubSub
	Logger *core.Logger
	Repo   repository.ServiceInventoryRepository
}

func NewServiceInventoryPubSub(ps *core.PubSub, rep repository.ServiceInventoryRepository, l *core.Logger) *ServiceInventoryPubSub {
	return &ServiceInventoryPubSub{
		PubSub: ps,
		Repo:   rep,
		Logger: l,
	}
}

func (serviceOrderPubSub *ServiceInventoryPubSub) RegisterSubscribers() {
	serviceOrderPubSub.subscriberServiceOrderStateChange()
}
