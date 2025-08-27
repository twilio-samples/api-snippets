package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			Record:                  "record-from-ringing-dual",
			RecordingStatusCallback: "www.myexample.com",
			InnerElements: []twiml.Element{
				&twiml.VoiceNumber{
					PhoneNumber: "+15558675310",
				},
			},
		},
	})

	fmt.Print(twiml)
}
