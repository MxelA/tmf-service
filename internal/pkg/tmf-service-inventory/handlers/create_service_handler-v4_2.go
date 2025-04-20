package service_inventory_handler_v42

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
)

func (h ServiceInventoryHandler) CreateServiceHandler(params service.CreateServiceParams) middleware.Responder {

	response := models.Service{}
	return service.NewCreateServiceCreated().WithPayload(&response)
}
