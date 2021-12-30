package service

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/william22913/learning/pion-janus-proxy/config"
	"github.com/william22913/learning/pion-janus-proxy/message"
)

func DoCreateSession(conn *websocket.Conn, config config.Configuration) (session message.SessionData, err error) {
	var response message.JanusResponse
	channel := make(chan message.JanusResponse)
	defer close(channel)

	go func() {
		err = conn.ReadJSON(&response)
		if err != nil {
			log.Println(err)
		}

		log.Println("response ->> ", response)

		channel <- response
	}()

	_, err = Write(conn, message.JanusRequest{
		Request: "create",
	})

	if err != nil {
		log.Println(err)
	}

	response = <-channel

	err = mapstructure.Decode(response.Data, &session)
	if err != nil {
		log.Println(err)
	}

	return
}

func Write(
	conn *websocket.Conn,
	request interface{},
) (string, error) {
	transactionID, err := uuid.NewRandomFromReader(rand.Reader)
	if err != nil {
		return "", err
	}

	transactionStr := transactionID.String()

	rawReq := make(map[string]interface{})

	err = mapstructure.Decode(request, &rawReq)
	if err != nil {
		return transactionStr, err
	}

	rawReq["transaction"] = transactionStr

	message, err := json.Marshal(rawReq)
	if err != nil {
		return transactionStr, err
	}

	fmt.Println(string(message))

	err = conn.WriteJSON(rawReq)
	if err != nil {
		return transactionStr, err
	}

	return transactionStr, nil
}
