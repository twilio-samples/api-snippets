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
    request := &twiliocomms.SendersSearchRequest{
        Address: "+14153902337",
        Channel: twiliocomms.SenderCommunicationChannelSms,
    }
    client.Senders.Search(
        context.TODO(),
        request,
    )
}
