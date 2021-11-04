package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/william22913/learning/nats/connect/create"
)

func main() {
	subject := "foo"
	conn, err := create.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	i := 1
	conn.Subscribe(subject, func(msg *nats.Msg) {
		i += 1
		printMsg(msg, i)
	})
	conn.Flush()

	if err := conn.LastError(); err != nil {
		fmt.Println(err)
	}

	for {
	}
}

func printMsg(m *nats.Msg, i int) {
	fmt.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
	fmt.Println()
}
