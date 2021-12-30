package service

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/william22913/learning/pion-janus-proxy/message"
)

func DoCreateRoom(conn *websocket.Conn, session message.SessionData, pluginHandle message.PluginHandle) (room message.RoomHandle, err error) {
	var response message.PluginResponse

	pluginChannel := make(chan message.PluginResponse)
	defer close(pluginChannel)

	go func() {

		err = conn.ReadJSON(&response)
		if err != nil {
			log.Println(err)
		}

		log.Println("response ->> ", response)

		pluginChannel <- response
	}()

	_, err = Write(conn, message.PluginRequest{
		Request: "message",
		Session: session.ID,
		Handle:  pluginHandle.ID,
		Body: message.TextRoomCreateRequest{
			Request: "create",
		},
	})

	if err != nil {
		log.Println(err)
	}

	response = <-pluginChannel

	room = message.RoomHandle{}

	err = mapstructure.Decode(
		response.Data.Data,
		&room,
	)

	return
}
