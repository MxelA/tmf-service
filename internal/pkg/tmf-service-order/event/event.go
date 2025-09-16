package event

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/pub_sub"
)

type EventFactory struct {
	tracer *core.Tracer
	logger *core.Logger
	pubSub *pub_sub.ServiceOrderPubSub
}

func NewEventFactory(pubSub *pub_sub.ServiceOrderPubSub, tracer *core.Tracer, logger *core.Logger) *EventFactory {
	return &EventFactory{
		pubSub: pubSub,
		tracer: tracer,
		logger: logger,
	}
}
