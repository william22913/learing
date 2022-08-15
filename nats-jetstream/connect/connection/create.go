package create

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/william22913/learning/nats/connect/create"
)

type AgoraServiceCode map[int]string

//Based on CallbackEvent from Agora
//https://docs.agora.io/en/cloud-recording/cloud_recording_callback_rest?platform=RESTful#event
func InitAgoraServiceCode() AgoraServiceCode {

	result := make(AgoraServiceCode)

	result[0] = "cloud.recording"
	result[1] = "recorder.module"
	result[2] = "uploader.module"
	result[4] = "extension"
	result[6] = "web.page"
	result[8] = "download"

	return result
}

func (asc AgoraServiceCode) ToArrayString(prefix string) []string {
	var result []string
	for key := range asc {
		result = append(result,
			fmt.Sprintf("%s.%s", prefix, asc[key]),
		)
	}

	return result
}

func CreateJetstreamConnection(
	a AgoraServiceCode,
) (conn *nats.Conn, ctx nats.JetStreamContext, err error) {

	conn, err = create.CreateConnection()
	if err != nil {
		return
	}

	ctx, err = conn.JetStream()
	if err != nil {
		return
	}

	stream, err := ctx.StreamInfo("agora-hook")
	if stream == nil {

		_, err = ctx.AddStream(&nats.StreamConfig{
			Name:      "agora-hook",
			Subjects:  a.ToArrayString("agora.hook"),
			Retention: nats.WorkQueuePolicy,
		})

	}

	if err != nil {
		return nil, ctx, err
	}

	return conn, ctx, err
}
