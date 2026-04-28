package example

import (
    context "context"

    twiliocomms "github.com/twilio/twilio-comms-go/twilio"
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
    request := []*twiliocomms.SenderPoolsAddSenderRequestItem{
        &twiliocomms.SenderPoolsAddSenderRequestItem{
            SenderId: "comms_sender_01h9krwprkeee8fzqspvwy6nq8",
        },
    }
    client.SenderPools.AddSender(
        context.TODO(),
        "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
        request,
    )
}
