package publisher

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/pub_sub"
)

type EventPublisher struct {
	tracer *core.Tracer
	logger *core.Logger
	pubSub *pub_sub.ServiceOrderPubSub
}

func NewEventPublisher(pubSub *pub_sub.ServiceOrderPubSub, tracer *core.Tracer, logger *core.Logger) *EventPublisher {
	return &EventPublisher{
		pubSub: pubSub,
		tracer: tracer,
		logger: logger,
	}
}
