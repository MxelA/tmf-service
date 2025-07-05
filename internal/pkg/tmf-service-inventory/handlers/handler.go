package handler

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
)

type ServiceInventoryHandler struct {
	repo   *repository.MongoServiceInventoryRepository
	logger *core.Logger
}

func NewServiceInventoryHandler(repo *repository.MongoServiceInventoryRepository, logger *core.Logger) *ServiceInventoryHandler {
	return &ServiceInventoryHandler{repo: repo, logger: logger}
}
