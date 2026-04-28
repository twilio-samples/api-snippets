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
    request := &twiliocomms.EmailsSendRequest{
        From: &twiliocomms.EmailAddressSender{
            Address: "support@example.company.io",
            Name: "Cool Co Support",
        },
        To: []*twiliocomms.EmailsSendRequestToItem{
            &twiliocomms.EmailsSendRequestToItem{
                EmailsSendRequestToItemAddress: &twiliocomms.EmailsSendRequestToItemAddress{
                    Address: "bob@example.com",
                    Name: twiliocomms.String(
                        "Bob Smith",
                    ),
                },
            },
        },
        Content: &twiliocomms.EmailHtmlContent{
            Html: "<html><body>Hey, <br/><br/><b>Cake</b></body></html>",
            Text: twiliocomms.String(
                "Hey, the cake is ready.",
            ),
            Subject: "Re: Wedding Cake",
            Attachments: &twiliocomms.EmailAttachments{
                &twiliocomms.EmailAttachmentsItem{
                    Filename: "filename",
                    ContentType: "contentType",
                    Content: "content",
                },
            },
        },
    }
    client.Emails.Send(
        context.TODO(),
        request,
    )
}
