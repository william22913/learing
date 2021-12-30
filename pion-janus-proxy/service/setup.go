package service

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/william22913/learning/pion-janus-proxy/message"
)

func DoSetup(conn *websocket.Conn, session message.SessionData, pluginHandle message.PluginHandle) (sdp string, err error) {
	channel := make(chan message.SetupResponse)
	defer close(channel)

	var response message.SetupResponse

	go func() {

		for {
			err = conn.ReadJSON(&response)
			if err != nil {
				log.Println(err)
			}

			log.Println("response ->> ", response)

			if response.Janus != "ack" {
				break
			}
		}

		channel <- response
	}()

	_, err = Write(conn, &message.PluginRequest{
		Request: "message",
		Session: session.ID,
		Handle:  pluginHandle.ID,
		Body: message.TextRoomCreateRequest{
			Request: "setup",
		},
	})

	if err != nil {
		log.Println(err)
	}

	response = <-channel

	sdp = response.Jsep.Sdp

	return
}
