package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
	connection "github.com/william22913/learning/nats-jetstream/connect/connection"
)

func main() {

	x := connection.InitAgoraServiceCode()

	conn, ctx, err := connection.CreateJetstreamConnection(x)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ctx.Subscribe("agora.hook.cloud.recording", func(msg *nats.Msg) {
		msg.Ack()
		fmt.Println(string(msg.Data))
	})

	// ctx.Publish("agora.hook.cloud-recording", []byte("1235551"))

	for {
	}

	// subj := "agora-hook:cloud-recording"
	// durableName := "agora-hook:cloud-recording"
	// consumer.AddConsumer(ctx, "agora-hook:cloud-recording", durableName, subj)

	// sub, err := ctx.PullSubscribe(subj, durableName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// go consumer.ReadSubs(sub)

	// sub, _ := ctx.PullSubscribe(subSubjectName, "order-review", nats.PullMaxWaiting(128))
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// for {
	// 	msgs, _ := sub.Fetch(10, nats.Context(ctx))
	// 	for _, msg := range msgs {
	// 		msg.Ack()
	// 		var order model.Order
	// 		err := json.Unmarshal(msg.Data, &order)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		log.Println("order-review service")
	// 		log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
	// 		reviewOrder(js, order)
	// 	}
	// }
}
