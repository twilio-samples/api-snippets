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
    request := &twiliocomms.MessagesSendRequest{
        To: []*twiliocomms.MessagesSendRequestToItem{
            &twiliocomms.MessagesSendRequestToItem{
                MessagesSendRequestToItemAddress: &twiliocomms.MessagesSendRequestToItemAddress{
                    Address: "+12065558844",
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
        From: &twiliocomms.MessagesSendRequestFrom{
            MessageAddressSender: &twiliocomms.MessageAddressSender{
                Address: "+14153901002",
                Channel: twiliocomms.MessageSenderChannelSms,
            },
        },
    }
    client.Messages.Send(
        context.TODO(),
        request,
    )
}
