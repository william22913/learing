package main

import (
	"fmt"

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

	err = conn.Publish(subject, []byte("test message"))
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Flush()
}
