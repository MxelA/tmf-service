package pub_sub

import (
	"fmt"
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	ServiceOrderStateChangeEventTopic string = "service_order_state_change"
)

func Init(ps *core.PubSub) {
	router := ps.GetRouter()
	router.AddNoPublisherHandler("test", ServiceOrderStateChangeEventTopic, ps.GetCore(), printMessages)
}

func printMessages(msg *message.Message) error {
	fmt.Printf(
		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}
