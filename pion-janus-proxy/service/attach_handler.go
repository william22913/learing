package service

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/william22913/learning/pion-janus-proxy/message"
)

func DoAttachHandler(conn *websocket.Conn, session message.SessionData) (pluginHandle message.PluginHandle, err error) {
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

	_, err = Write(conn, &message.SessionAttachRequest{
		Request: "attach",
		Session: session.ID,
		Plugin:  "janus.plugin.textroom",
		Tag:     "transport",
	})

	if err != nil {
		log.Println(err)
	}

	response = <-channel

	err = mapstructure.Decode(response.Data, &pluginHandle)

	if err != nil {
		log.Println(err)
	}

	return
}
