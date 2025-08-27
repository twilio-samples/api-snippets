package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceConference{
					Name:                    "LoveTwilio",
					Record:                  "record-from-start",
					RecordingStatusCallback: "www.myexample.com",
				},
			},
		},
	})

	fmt.Print(twiml)
}
