package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
	"github.com/pion/webrtc/v3"
	"github.com/william22913/learning/pion-janus-proxy/config"
	"github.com/william22913/learning/pion-janus-proxy/service"
)

func main() {

	conn, err := Connect(
		Default.Janus.Url,
	)

	if err != nil {
		log.Println("1", err)
		return
	}

	// err = service.DoHandshake(conn, Default)
	// if err != nil {
	// 	log.Println("2", err)
	// 	return
	// }

	session, err := service.DoCreateSession(conn, Default)
	if err != nil {
		log.Println("3", err)
		return
	}

	handler, err := service.DoAttachHandler(conn, session)
	if err != nil {
		log.Println("4", err)
		return
	}

	room, err := service.DoCreateRoom(conn, session, handler)
	if err != nil {
		log.Println("5", err)
		return
	}

	fmt.Println(room)

	sdp, err := service.DoSetup(conn, session, handler)
	if err != nil {
		log.Println("6", err)
		return
	}

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
		// SDPSemantics: webrtc.SDPSemanticsUnifiedPlanWithFallback,
	}

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Println("7", err)
		return
	}

	defer func() {
		if err := peerConnection.Close(); err != nil {
			log.Println("8", err)
		}
	}()

	peerConnection.OnConnectionStateChange(func(pcs webrtc.PeerConnectionState) {
		fmt.Printf("Connection State has changed %s \n", pcs.String())

		if pcs == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity
			// for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are
			// interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			log.Println("Peer Connection has gone to failed exiting")
			os.Exit(0)
		}
	})

	option := &webrtc.DataChannelInit{}

	channel, err := peerConnection.CreateDataChannel("JanusDataChannel", option)
	if err != nil {
		log.Println("9", err)
	}

	channel.OnOpen(func() {
		fmt.Println("datachannel open")
	})

	channel.OnClose(func() {
		fmt.Println("datachannel close")
	})

	// Add handlers for setting up the connection.
	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Println(state)
	})

	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		err = service.DoTrickle(conn, session, handler, candidate)
		if err != nil {
			log.Println("10", err)
		}
	})

	peerConnection.OnDataChannel(func(dc *webrtc.DataChannel) {

		fmt.Println(dc.Label(), *dc.ID(), "New data channel")

		// Register channel opening handling
		dc.OnOpen(func() {
			fmt.Println("open")
			// dataChannels[dc.Label()] = dc
			// dc ->
		})

		dc.OnClose(func() {
			// label := dc.Label()
			//
			// -> reconnect
			// delete(dataChannels, label)
		})

		dc.OnError(func(err error) {
			//
			// -> peer connection re-negotiate

		})

		dc.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Println(string(msg.Data))
			// <-
		})
	})

	sdpDesc := webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  sdp,
	}

	peerConnection.SetRemoteDescription(sdpDesc)
	if err != nil {
		log.Println("11", err)
		return
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		log.Println("12", err)
		return
	}

	fmt.Println("aswer", answer)

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		log.Println("13", err)
		return
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	time.Sleep(5 * time.Second)

	err = service.DoPluginAck(conn, session, handler, sdp)
	if err != nil {
		log.Println("14", err)
		return
	}

	join := fmt.Sprintf(`
		{
			"textroom": "join",
			"room": %s,
			"username": "wills",
			"transaction": "hehe12487h"
		}
	`, room.Room)

	err = channel.SendText(join)
	if err != nil {
		log.Println("15", err)
		return
	}

	message := fmt.Sprintf(`
		{
			"textroom": "message",
			"room": %s,
			"text": "hello world",
			"transaction": "hehe12487h"
		}
	`, room.Room)

	err = channel.SendText(message)
	if err != nil {
		log.Println("16", err)
		return
	}

	for {
	}

}

var (
	Default config.Configuration
)

func init() {

	err := envconfig.Process("JGPROXY", &Default)

	if err != nil {
		log.Println(err)
	}

}

type KeepAliveMessage struct {
	Request     string `json:"janus"`
	Session     string `json:"session_id"`
	Transaction string `json:"transaction"`
}

func Connect(url string) (*websocket.Conn, error) {
	upgradeRequest := http.Header{}
	// upgradeRequest.Add("Connection", "Upgrade")
	// upgradeRequest.Add("Upgrade", "websocket")
	// Note: add janus-protocol to WebSocket protocol
	upgradeRequest.Add("Sec-Websocket-Protocol", "janus-protocol")

	dealer := websocket.DefaultDialer
	// upgradeRequest.
	conn, _, err := dealer.Dial(url, upgradeRequest)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func RoomId() (string, error) {
	id, err := uuid.NewRandomFromReader(rand.Reader)

	return id.String(), err
}
