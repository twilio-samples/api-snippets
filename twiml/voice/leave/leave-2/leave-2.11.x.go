package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePlay{
			Url: "http://com.twilio.sounds.music.s3.amazonaws.com/MARKOVICHAMP-Borghestral.mp3",
		},
	})

	fmt.Print(twiml)
}
