package service

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/william22913/learning/pion-janus-proxy/config"
	"github.com/william22913/learning/pion-janus-proxy/message"
)

func DoHandshake(conn *websocket.Conn, config config.Configuration) (err error) {

	channel := make(chan message.JanusResponse)
	defer close(channel)

	var response message.JanusResponse
	go func() {

		err = conn.ReadJSON(&response)
		if err != nil {
			log.Println(err)
		}

		log.Println("response ->> ", response)

		channel <- response
	}()

	err = Handshake(
		conn,
		config.Janus.Conference,
		message.UserToken{},
		"transport",
		channel,
	)

	return err

}

func Handshake(
	conn *websocket.Conn,
	conference string,
	token message.UserToken,
	tag string,
	channel chan message.JanusResponse,
) error {

	err := conn.WriteJSON(message.HandshakeRequest{
		UserId:      "729QGoN41A",
		Role:        "admin",
		Conference:  "d0ccd4b1-7b64-49e9-9c41-36ab75ced8a0",
		Transaction: "1dfihfasasudasdbkasjdas",
		Token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJydCI6ImV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpqYVdRaU9pSTNNamxSUjI5T05ERkJJaXdpWlhod0lqb3hOalF3TkRBMU5USXhMQ0oxYVdRaU9qRXpOVEkxTmpFeE15d2lkVzlqSWpvaVNrbE1URmxWT1VjM09WQlZXRTlJU3lJc0luSWlPaUpoWkcxcGJpSXNJbVJwWkNJNkltbGtJaXdpWkc0aU9pSnVZVzFsSWl3aWRHOXJaVzVKUkNJNklqRTJOREF6TVRrd05qRTVOelV4TVRFM01qa2lmUS5oNWtBaDRYdURiMGVQSlpfd0VORjctTF95Zmh6d3Q1TVZWbVhuTUlSSDQ0IiwiZXhwIjoxNjQwMzE5OTYxLCJ1aWQiOjEzNTI1NjExMywidW9jIjoiSklMTFlVOUc3OVBVWE9ISyIsInIiOiJhZG1pbiIsImRpZCI6ImlkIiwiZG4iOiJuYW1lIiwidG9rZW5JRCI6IjE2NDAzMTkwNjE5NzUxMTE3MjkifQ.x7lIkIRs6wx342RDaoKwgmvBqLvQKub24x3ld64tCxU",
	})

	if err != nil {
		return err
	}

	response := <-channel

	if response.Status != "authorized" {
		return fmt.Errorf("not authorized")
	}

	return nil
}
