package example

import (
    context "context"

    client "github.com/twilio/twilio-comms-go/twilio/client"
    option "github.com/twilio/twilio-comms-go/twilio/option"
)

func do() {
    client := client.NewWithOptions(
        option.WithBasicAuth(
            "TWILIO_ACCOUNT_SID",
            "TWILIO_AUTH_TOKEN",
        ),
    )
    client.SenderPools.RemoveSender(
        context.TODO(),
        "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
        "comms_sender_01h9krwprkeee8fzqspvwy6nq8",
    )
}
