package pub_sub

import (
	"fmt"
)

func (serviceOrderInventoryPubSub *ServiceInventoryPubSub) ServiceOrderAttributeValueChangeSubscriber() {
	router := serviceOrderInventoryPubSub.pubSub.GetRouter()
	router.AddNoPublisherHandler(fmt.Sprintf("service-inventory_%s", ServiceOrderAttributeValueChangeEventTopic), ServiceOrderAttributeValueChangeEventTopic, serviceOrderInventoryPubSub.pubSub.GetCore(), serviceOrderInventoryPubSub.eventSubscriber.ServiceOrderChangeSubscriberHandler)
}
