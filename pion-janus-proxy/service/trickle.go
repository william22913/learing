package service

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"github.com/william22913/learning/pion-janus-proxy/message"
)

func DoTrickle(conn *websocket.Conn, session message.SessionData, pluginHandle message.PluginHandle, candidate *webrtc.ICECandidate) (err error) {
	channel := make(chan message.SetupResponse)
	defer close(channel)

	var response message.SetupResponse

	go func() {

		err = conn.ReadJSON(&response)
		if err != nil {
			log.Println(err)
		}

		log.Println("response ->> ", response)

		channel <- response
	}()

	msg := message.TrickleRequest{
		Request: "trickle",
		Session: session.ID,
		Handle:  pluginHandle.ID,
	}

	if candidate != nil {
		msg.Candidate = candidate
	} else {
		msg.Candidate = message.TrickleRequestComplete{
			Completed: true,
		}
	}

	_, err = Write(conn, &msg)

	if err != nil {
		log.Println(err)
	}

	<-channel

	return
}
