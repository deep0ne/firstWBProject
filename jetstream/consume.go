package jetstream

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/deep0ne/firstWBProject/models"
	"github.com/nats-io/nats.go"
)

const (
	subjectNameConsume = "ORDERS.created"
)

func ConsumeOrders(js nats.JetStreamContext) ([]models.Order, error) {
	sub, _ := js.PullSubscribe(subjectNameConsume, "order-review", nats.PullMaxWaiting(128))
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	orders := make([]models.Order, 0)
	for {
		select {
		case <-ctx.Done():
			return orders, nil
		default:
		}
		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			msg.Ack()
			var order models.Order
			err := json.Unmarshal(msg.Data, &order)
			if err != nil {
				return orders, err
			}
			orders = append(orders, order)
			log.Printf("Order with UID %s is consumed and is going to postgres...\n", order.OrderInfo.OrderUID)
		}
	}
}
