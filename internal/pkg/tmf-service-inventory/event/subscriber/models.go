package subscriber

import "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"

// Root event wrapper
type EventWrapper struct {
	Event Event `json:"event"`
}

// Event object
type Event struct {
	ServiceOrder ServiceOrder `json:"serviceOrder"`
}

// ServiceOrder object
type ServiceOrder struct {
	ID               string             `json:"id"`
	State            string             `json:"state"`
	Href             string             `json:"href"`
	ServiceOrderItem []ServiceOrderItem `json:"serviceOrderItem"`
}

// Single ServiceOrderItem
type ServiceOrderItem struct {
	ID      string                     `json:"id"`
	Action  models.OrderItemActionType `json:"action"`
	State   string                     `json:"state"`
	Service models.Service             `json:"service"`
}
