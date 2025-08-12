package handler

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/repositories"
)

type ServiceOrderHandler struct {
	repo   repository.ServiceOrderRepository
	logger *core.Logger
}

func NewServiceOrderHandler(repo repository.ServiceOrderRepository, logger *core.Logger) *ServiceOrderHandler {
	return &ServiceOrderHandler{repo: repo, logger: logger}
}
