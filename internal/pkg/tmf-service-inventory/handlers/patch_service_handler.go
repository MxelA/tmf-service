package handler

import (
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
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

	var jpo []*models.JSONPatchOperation
	
	//marshal to json bytes
	raw, err := json.Marshal(req.Service)
	if err != nil {
		errCode := "400"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	// Unmarshal to JSONPatchOperation struct
	if err := json.Unmarshal(raw, &jpo); err != nil {
		errCode := "400"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	// Validate JSONPatchOperation struct
	for _, op := range jpo {
		if err := op.Validate(strfmt.Default); err != nil {
			errCode := "422"
			reason := "Validation error"
			var errModel = models.Error{
				Code:    &errCode,
				Message: err.Error(),
				Reason:  &reason,
			}
			h.logger.GetCore().Error(err.Error())
			return service.NewPatchServiceBadRequest().WithPayload(&errModel)
		}
	}
	//if err := jpo.Validate(strfmt.Default); err != nil {
	//	errCode := "606"
	//	reason := "Validation error"
	//	var errModel = models.Error{
	//		Code:    &errCode,
	//		Message: err.Error(),
	//		Reason:  &reason,
	//	}
	//	log.Println(err)
	//	return service.NewPatchServiceBadRequest().WithPayload(&errModel)
	//}

	// make patch operation
	patchOperation, err := jsonpatch.DecodePatch(raw)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	// get service from DB
	serviceEntity, err := h.repo.GetByID(req.HTTPRequest.Context(), req.ID, nil)
	if err != nil {
		errCode := "404"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Service not found",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceNotFound().WithPayload(&errModel)
	}

	serviceJson, err := json.Marshal(serviceEntity)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}
	modifiedServiceJSON, err := patchOperation.Apply(serviceJson)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	var modifiedServiceEntity models.Service
	if err := json.Unmarshal(modifiedServiceJSON, &modifiedServiceEntity); err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	_, err = h.repo.Update(req.HTTPRequest.Context(), req.ID, &modifiedServiceEntity)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	return service.NewPatchServiceOK().WithPayload(&modifiedServiceEntity)
}

func processMergePatch(h *ServiceInventoryHandler, req service.PatchServiceParams) middleware.Responder {

	var serviceUpdate models.ServiceUpdate

	//marshal to json bytes
	raw, err := json.Marshal(req.Service)
	if err != nil {
		errCode := "400"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	// Unmarshal to struct
	if err = json.Unmarshal(raw, &serviceUpdate); err != nil {
		errCode := "500"
		reason := "Internal server error"
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: err.Error(),
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	// Validate serviceUpdate struct
	if err := serviceUpdate.Validate(strfmt.Default); err != nil {
		errCode := "606"
		reason := "Validation error"
		var errModel = models.Error{
			Code:    &errCode,
			Message: err.Error(),
			Reason:  &reason,
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceBadRequest().WithPayload(&errModel)
	}

	// Merge patch
	_, err = h.repo.Update(req.HTTPRequest.Context(), req.ID, &serviceUpdate)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service.NewPatchServiceInternalServerError().WithPayload(&errModel)
	}

	retrieveServiceOrder, err := h.repo.GetByID(req.HTTPRequest.Context(), req.ID, nil)

	return service.NewPatchServiceOK().WithPayload(retrieveServiceOrder)
}
