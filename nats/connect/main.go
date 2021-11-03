package main

import (
	"fmt"

	"github.com/william22913/learning/nats/connect/create"
)

func main() {
	conn, err := create.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conn)
}
