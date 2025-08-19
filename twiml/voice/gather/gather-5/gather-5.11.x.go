package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceGather{
			Action: "/process_gather.php",
			Method: "GET",
			InnerElements: []twiml.Element{
				&twiml.VoiceSay{
					Message: "Enter something, or not",
				},
			},
		},
		&twiml.VoiceRedirect{
			Method: "GET",
			Url:    "/process_gather.php?Digits=TIMEOUT",
		},
	})

	fmt.Print(twiml)
}
