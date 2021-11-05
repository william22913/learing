package create

import (
	"flag"

	"github.com/nats-io/nats.go"
)

func CreateConnection() (conn *nats.Conn, err error) {
	var urls = flag.String("s", "nats://127.0.0.1:4444", "The nats server URLs (separated by comma)")
	opts := []nats.Option{nats.Name("NATS Sample Publisher"), nats.UseOldRequestStyle()}

	conn, err = nats.Connect(*urls, opts...)
	if err != nil {
		return
	}

	return conn, nil
}
