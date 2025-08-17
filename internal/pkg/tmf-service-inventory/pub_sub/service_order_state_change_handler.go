package pub_sub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/jinzhu/copier"
)

func (serviceOrderPubSub *ServiceInventoryPubSub) subscriberServiceOrderStateChange() {
	router := serviceOrderPubSub.PubSub.GetRouter()
	router.AddNoPublisherHandler(fmt.Sprintf("servcie-inventory_%s", ServiceOrderStateChangeEventTopic), ServiceOrderStateChangeEventTopic, serviceOrderPubSub.PubSub.GetCore(), serviceOrderPubSub.subscriberServiceOrderStateChangeHandler)
}

func (serviceOrderPubSub *ServiceInventoryPubSub) subscriberServiceOrderStateChangeHandler(msg *message.Message) error {
	serviceOrderPubSub.Logger.GetCore().Debug("\n> Received message: %s\n> %s\n> metadata: %v", msg.UUID, string(msg.Payload), msg.Metadata)
	//fmt.Printf(
	//	"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
	//	msg.UUID, string(msg.Payload), msg.Metadata,
	//)

	services, err := parseServiceOrderStateChangeMsg(msg)
	if err != nil {
		return err
	}

	for _, service := range services {
		if service.ID != nil {
			_, err := serviceOrderPubSub.Repo.Update(context.Background(), *service.ID, service)
			if err != nil {
				return err
			}
		} else {
			serviceCreate := MapToServiceCreate(*service)
			_, _ = serviceOrderPubSub.Repo.Create(context.Background(), &serviceCreate)
		}

	}

	return nil
}

func parseServiceOrderStateChangeMsg(msg *message.Message) ([]*models.Service, error) {
	var services []*models.Service
	var evt struct {
		Event struct {
			ServiceOrder struct {
				ID               string `json:"id"`
				State            string `json:"state"`
				Href             string `json:"href"`
				ServiceOrderItem []struct {
					ID      string         `json:"id"`
					Action  string         `json:"action"`
					State   string         `json:"state"`
					Service models.Service `json:"service"`
				} `json:"serviceOrderItem"`
			} `json:"serviceOrder"`
		} `json:"event"`
	}

	if err := json.Unmarshal(msg.Payload, &evt); err != nil {
		return services, err
	}

	for _, item := range evt.Event.ServiceOrder.ServiceOrderItem {
		if item.State == "completed" {
			service := item.Service
			service.ServiceOrderItem = []*models.RelatedServiceOrderItem{
				{
					ServiceOrderID:   &evt.Event.ServiceOrder.ID,
					ItemID:           item.ID,
					ServiceOrderHref: &evt.Event.ServiceOrder.Href,
					ItemAction:       models.OrderItemActionType(item.Action),
				},
			}
			services = append(services, &service)
		}
	}

	return services, nil
}

func MapToServiceCreate(svc models.Service) models.ServiceCreate {
	var create models.ServiceCreate
	_ = copier.Copy(&create, &svc)
	return create
}

func MapToServiceUpdate(svc models.Service) models.ServiceUpdate {
	var update models.ServiceUpdate
	_ = copier.Copy(&update, &svc)
	return update
}
