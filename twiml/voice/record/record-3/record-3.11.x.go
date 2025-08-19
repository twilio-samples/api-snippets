package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceSay{
			Message: `
	Please leave a message at the beep. 
	Press the star key when finished.
`,
		},
		&twiml.VoiceRecord{
			Action:      "http://foo.edu/handleRecording.php",
			FinishOnKey: "*",
			MaxLength:   "20",
			Method:      "GET",
		},
		&twiml.VoiceSay{
			Message: "I did not receive a recording",
		},
	})

	fmt.Print(twiml)
}
