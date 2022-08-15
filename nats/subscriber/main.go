package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/william22913/learning/nats/connect/create"
)

func main() {
	subject := "agora"
	conn, err := create.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	i := 1
	_, err = conn.Subscribe(subject, func(msg *nats.Msg) {
		i += 1
		printMsg(msg, i)
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create Conference -> rooms di many janus gateway proxy
	// Subscribe "conference id"
	// message -> {  }

	//for ping-pong
	err = conn.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

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
