package pub_sub

import (
	"github.com/MxelA/tmf-service/internal/core"
	repository "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
)

const (
	ServiceOrderAttributeValueChangeEventTopic string = "service_order_attribute_value_change"
)

type ServiceInventoryPubSub struct {
	PubSub *core.PubSub
	Logger *core.Logger
	Repo   repository.ServiceInventoryRepository
	Tracer *core.Tracer
}

func NewServiceInventoryPubSub(ps *core.PubSub, rep repository.ServiceInventoryRepository, l *core.Logger, tr *core.Tracer) *ServiceInventoryPubSub {
	return &ServiceInventoryPubSub{
		PubSub: ps,
		Repo:   rep,
		Logger: l,
		Tracer: tr,
	}
}

func (serviceOrderInventoryPubSub *ServiceInventoryPubSub) RegisterSubscribers() {
	serviceOrderInventoryPubSub.ServiceOrderAttributeValueChangeSubscriber()
}
