package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"strings"
)

func (h *ServiceInventoryHandler) PatchServiceHandler(req service.PatchServiceParams) middleware.Responder {
	contentType := req.HTTPRequest.Header["Content-Type"][0]

	if strings.Contains(contentType, "application/json-patch+json") {
		return processJsonPatch(h, req)
	} else if strings.Contains(contentType, "application/merge-patch+json") {
		return processMergePatch(h, req)
	}

	errCode := "422"
	reason := "Unsupported Media Type"
	errModel := models.Error{
		Reason:  &reason,
		Code:    &errCode,
		Message: "Unsupported media type " + contentType,
	}

	return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
}

func processJsonPatch(h *ServiceInventoryHandler, req service.PatchServiceParams) middleware.Responder {
	//TO DO: Add JSON Patch support
	errCode := "422"
	reason := "Unsupported media type "
	var errModel = models.Error{
		Reason:  &reason,
		Code:    &errCode,
		Message: "Unsupported 'application/json-patch+json' media type ",
	}
	return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
}

func processMergePatch(h *ServiceInventoryHandler, req service.PatchServiceParams) middleware.Responder {

	_, err := h.repo.MergePatch(req.HTTPRequest.Context(), req.ID, req.Service)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	retrieveServiceOrder, err := h.repo.GetByID(req.HTTPRequest.Context(), req.ID, nil)

	return service.NewPatchServiceOK().WithPayload(retrieveServiceOrder)
}
