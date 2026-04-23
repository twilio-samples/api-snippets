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
    request := &twiliocomms.MessagesSendRequest{
        To: []*twiliocomms.MessagesSendRequestToItem{
            &twiliocomms.MessagesSendRequestToItem{
                MessagesSendRequestToItemAddress: &twiliocomms.MessagesSendRequestToItemAddress{
                    Address: "+14153902337",
                    Channel: twiliocomms.MessageAddressChannelPhone,
                },
            },
        },
        Content: &twiliocomms.MessagesSendRequestContent{
            MessageContentTextWithMedia: &twiliocomms.MessageContentTextWithMedia{
                Text: twiliocomms.String(
                    "Hello, World!",
                ),
            },
        },
    }
    client.Messages.Send(
        context.TODO(),
        request,
    )
}
