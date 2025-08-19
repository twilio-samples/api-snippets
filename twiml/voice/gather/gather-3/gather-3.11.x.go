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
					Message: "Please enter your account number, followed by the pound sign",
				},
			},
		},
		&twiml.VoiceSay{
			Message: "We didn't receive any input. Goodbye!",
		},
	})

	fmt.Print(twiml)
}
