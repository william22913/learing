package create

import (
	"github.com/nats-io/nats.go"
	"github.com/william22913/learning/nats/connect/create"
)

func CreateJetstreamConnection() (conn *nats.Conn, ctx nats.JetStreamContext, err error) {
	conn, err = create.CreateConnection()
	if err != nil {
		return
	}

	ctx, err = conn.JetStream()
	if err != nil {
		return
	}

	return conn, ctx, err
}
