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
            "<username>",
            "<password>",
        ),
    )
    request := &twiliocomms.SenderPoolsUpdateRequest{
        Name: twiliocomms.String(
            "Customer Support Senders - Tier 1",
        ),
        Tags: &twiliocomms.Tags{
            "tier": "1",
        },
    }
    client.SenderPools.Update(
        context.TODO(),
        "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
        request,
    )
}
