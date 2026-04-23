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
    request := &twiliocomms.SendersResolveRequest{
        SendersResolveRequestSenderPoolId: &twiliocomms.SendersResolveRequestSenderPoolId{
            RecipientAddresses: []*twiliocomms.CommunicationRecipient{
                &twiliocomms.CommunicationRecipient{
                    Address: "+14153902337",
                    Channel: twiliocomms.CommunicationChannelPhone,
                },
                &twiliocomms.CommunicationRecipient{
                    Address: "+14153902337",
                    Channel: twiliocomms.CommunicationChannelWhatsapp,
                },
                &twiliocomms.CommunicationRecipient{
                    Address: "davidpletnjov@example.com",
                    Channel: twiliocomms.CommunicationChannelEmail,
                },
            },
        },
    }
    client.Senders.Resolve(
        context.TODO(),
        request,
    )
}
