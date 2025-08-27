package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceConnect{
			Action: "https://myhttpserver.com/connect_action",
			InnerElements: []twiml.Element{
				&twiml.VoiceConversationRelay{
					Url:             "wss://mywebsocketserver.com/websocket",
					WelcomeGreeting: "Hi! Ask me anything!",
				},
			},
		},
	})

	fmt.Print(twiml)
}
