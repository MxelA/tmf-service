package subscriber

import (
	"github.com/MxelA/tmf-service/internal/core"
	repository "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
)

type EventSubscriber struct {
	tracer *core.Tracer
	logger *core.Logger
	repo   repository.ServiceInventoryRepository
}

func NewEventSubscriber(repo repository.ServiceInventoryRepository, tracer *core.Tracer, logger *core.Logger) *EventSubscriber {
	return &EventSubscriber{
		repo:   repo,
		tracer: tracer,
		logger: logger,
	}
}
