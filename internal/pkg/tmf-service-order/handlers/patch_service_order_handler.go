package handler

import (
	"encoding/json"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service/internal/utils"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func (h *ServiceOrderHandler) PatchServiceOrderHandler(req service_order.PatchServiceOrderParams) middleware.Responder {

	patchType := utils.DetectPatchMediaType(req.HTTPRequest.Header)
	if *patchType == utils.JSONPatch {
		return processJsonPatch(h, req)
	} else if *patchType == utils.MergePatch {
		return processMergePatch(h, req)
	}

	errCode := "422"
	reason := "Unsupported Media Type"
	errModel := models.Error{
		Reason:  &reason,
		Code:    &errCode,
		Message: "Unsupported media type ",
	}

	return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
}

func processJsonPatch(h *ServiceOrderHandler, req service_order.PatchServiceOrderParams) middleware.Responder {

	var jpo []*models.JSONPatchOperation

	//marshal to json bytes
	raw, err := json.Marshal(req.ServiceOrder)
	if err != nil {
		errCode := "400"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
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
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
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
			return service_order.NewPatchServiceOrderBadRequest().WithPayload(&errModel)
		}
	}

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
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}

	// get service from DB
	serviceOrderEntity, err := h.repo.GetByID(req.HTTPRequest.Context(), req.ID, nil)
	if err != nil {
		errCode := "404"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Service not found",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderNotFound().WithPayload(&errModel)
	}

	serviceOrderJson, err := json.Marshal(serviceOrderEntity)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}
	modifiedServiceOrderJSON, err := patchOperation.Apply(serviceOrderJson)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}

	var modifiedServiceOrderEntity models.ServiceOrder
	if err := json.Unmarshal(modifiedServiceOrderJSON, &modifiedServiceOrderEntity); err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}

	_, err = h.repo.Update(req.HTTPRequest.Context(), req.ID, &modifiedServiceOrderEntity)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}

	return service_order.NewPatchServiceOrderOK().WithPayload(&modifiedServiceOrderEntity)
}

func processMergePatch(h *ServiceOrderHandler, req service_order.PatchServiceOrderParams) middleware.Responder {

	var serviceOrderUpdate models.ServiceOrderUpdate

	//marshal to json bytes
	raw, err := json.Marshal(req.ServiceOrder)
	if err != nil {
		errCode := "400"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}

	// Unmarshal to struct
	if err = json.Unmarshal(raw, &serviceOrderUpdate); err != nil {
		errCode := "500"
		reason := "Internal server error"
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: err.Error(),
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}

	// Validate serviceUpdate struct
	if err := serviceOrderUpdate.Validate(strfmt.Default); err != nil {
		errCode := "606"
		reason := "Validation error"
		var errModel = models.Error{
			Code:    &errCode,
			Message: err.Error(),
			Reason:  &reason,
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderBadRequest().WithPayload(&errModel)
	}

	// Merge patch
	_, err = h.repo.Update(req.HTTPRequest.Context(), req.ID, &serviceOrderUpdate)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		h.logger.GetCore().Error(err.Error())
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}

	retrieveServiceOrder, err := h.repo.GetByID(req.HTTPRequest.Context(), req.ID, nil)

	return service_order.NewPatchServiceOrderOK().WithPayload(retrieveServiceOrder)
}
