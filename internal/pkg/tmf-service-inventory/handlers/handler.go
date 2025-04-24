package service_inventory_handler_v42

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
)

type ServiceInventoryHandler struct {
	repo service_inventory_repository.ServiceInventoryRepository
}

func NewServiceInventoryHandler(repo service_inventory_repository.ServiceInventoryRepository) *ServiceInventoryHandler {
	return &ServiceInventoryHandler{repo: repo}
}
