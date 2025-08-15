package pub_sub

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
)

func (serviceOrderPubSub *ServiceInventoryPubSub) RegisterServiceOrderStateChangeSubscriber() {
	router := serviceOrderPubSub.PubSub.GetRouter()
	router.AddNoPublisherHandler(fmt.Sprintf("servcie-inventory_%s", ServiceOrderStateChangeEventTopic), ServiceOrderStateChangeEventTopic, serviceOrderPubSub.PubSub.GetCore(), func(msg *message.Message) error {
		fmt.Printf(
			"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
			msg.UUID, string(msg.Payload), msg.Metadata,
		)
		return nil
	})
}
