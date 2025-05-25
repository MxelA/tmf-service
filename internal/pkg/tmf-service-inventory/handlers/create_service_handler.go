package handler

import (
	"context"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_0/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_0/server/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
)

func (h ServiceInventoryHandler) CreateServiceHandler(params service.CreateServiceParams) middleware.Responder {
	_ = h.repo.Create(context.Background(), params.Service)
	response := models.Service{}
	return service.NewCreateServiceCreated().WithPayload(&response)
}
