package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func (h ServiceInventoryHandler) CreateServiceHandler(params service.CreateServiceParams) middleware.Responder {
	insertResult, err := h.repo.Create(params.HTTPRequest.Context(), params.Service)
	//
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}

		return service.NewCreateServiceInternalServerError().WithPayload(&errModel)
	}

	// Get mongo document
	id := insertResult.InsertedID.(primitive.ObjectID).Hex()

	retrieveService, err := h.repo.GetByID(params.HTTPRequest.Context(), id, nil)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		errModel := models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service.NewRetrieveServiceInternalServerError().WithPayload(&errModel)
	}

	return service.NewCreateServiceCreated().WithPayload(retrieveService)
}
