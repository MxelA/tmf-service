package handler

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-order/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
)

func (h *ServiceOrderHandler) ListServiceOrderHandler(params service_order.ListServiceOrderParams) middleware.Responder {

	retrieveServiceOrders, retrieveServiceOrdersTotalCount, err := h.repo.GetAllPaginate(params.HTTPRequest.Context(), params.HTTPRequest, params.Fields, params.Offset, params.Limit)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}

		h.logger.GetCore().Error(err.Error())

		return service_order.NewListServiceOrderInternalServerError().WithPayload(&errModel)
	}
	return service_order.NewListServiceOrderOK().WithPayload(retrieveServiceOrders).
		WithXTotalCount(*retrieveServiceOrdersTotalCount).
		WithXResultCount(int64(len(retrieveServiceOrders))).
		WithPayload(retrieveServiceOrders)
}
