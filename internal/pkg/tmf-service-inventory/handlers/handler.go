package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
)

type ServiceInventoryHandler struct {
	repo repository.ServiceInventoryRepository
}

func NewServiceInventoryHandler(repo repository.ServiceInventoryRepository) *ServiceInventoryHandler {
	return &ServiceInventoryHandler{repo: repo}
}
