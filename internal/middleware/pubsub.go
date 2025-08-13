package middleware

import (
	"github.com/MxelA/tmf-service/internal/core"
	"github.com/ThreeDotsLabs/watermill/message"
)

// Declare and register Pub/Sub middleware, make sure it returns message.HandlerFunc
// type otherwise it won't be added to router stack.

func InitPubSubMiddleware(ps *core.PubSub, mws ...message.HandlerMiddleware) {
	router := ps.GetRouter()

	for _, mw := range mws {
		router.AddMiddleware(mw)
	}
}
