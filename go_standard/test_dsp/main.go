package main

import (
	"github.com/pixelbender/go-sdp/sdp"
	"fmt"
)

func main() {
	sess := &sdp.Session{
		Origin: &sdp.Origin{
			Username:       "alice",
			Address:        "alice.example.org",
			SessionID:      2890844526,
			SessionVersion: 2890844526,
		},
		Name:       "Example",
		Connection: &sdp.Connection{
			Address: "127.0.0.1",
		},
		Media: []*sdp.Media{
			{
				Type:  "audio",
				Port:  10000,
				Proto: "RTP/AVP",
				Formats: []*sdp.Format{
					{Payload: 0, Name: "PCMU", ClockRate: 8000},
					{Payload: 8, Name: "PCMA", ClockRate: 8000},
				},
			},
		},
		Mode: sdp.ModeSendRecv,
	}

	fmt.Println(sess.String())
}
