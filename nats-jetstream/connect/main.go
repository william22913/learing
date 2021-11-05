package main

import (
	"fmt"

	connection "github.com/william22913/learning/nats-jetstream/connect/connection"
	consumer "github.com/william22913/learning/nats-jetstream/consumer"
)

func main() {
	conn, ctx, err := connection.CreateJetstreamConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	subj := "foo"
	durableName := "foo_test"
	consumer.AddConsumer(ctx, "FOO", durableName, subj)

	sub, err := ctx.PullSubscribe(subj, durableName)
	if err != nil {
		fmt.Println(err)
		return
	}

	consumer.ReadSubs(sub)

	for {

	}
}
