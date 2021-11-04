package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func AddConsumer(js nats.JetStreamContext, streamName, durableName, subject string) (info *nats.ConsumerInfo, err error) {
	info, err = js.AddConsumer(streamName, &nats.ConsumerConfig{
		Durable:   durableName,
		AckPolicy: nats.AckAllPolicy,
		// MaxAckPending: 1,      // default value is 20,000
		FilterSubject: subject,
	})
	if err != nil {
		return
	}

	return info, err
}

func ReadSubs(sub *nats.Subscription) (err error) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		msgs, err := sub.Fetch(10, nats.Context(ctx))
		if err != nil && err.Error() != "context deadline exceeded" {
			fmt.Println(err)
			return
		}

		for i := 0; i < len(msgs); i++ {
			msgs[i].Ack()
			var consInfo *nats.ConsumerInfo
			consInfo, err = sub.ConsumerInfo()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%s, %s, %s", consInfo.Stream, consInfo.Name, msgs[i].Data)
			fmt.Println()
		}
	}()

	return ReadSubs(sub)
}
