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
    request := &twiliocomms.PushNotificationsSendRequest{
        From: &twiliocomms.PushNotificationsSendRequestFrom{
            PushNotificationCredentialSender: &twiliocomms.PushNotificationCredentialSender{
                Fcm: twiliocomms.String(
                    "comms_credential_01h9krwprkeee8fzqspvwy6nq8",
                ),
                Apn: twiliocomms.String(
                    "comms_credential_8qn6ywvpsqzf8eeekrpwrk9h10",
                ),
            },
        },
        To: []*twiliocomms.PushNotificationsSendRequestToItem{
            &twiliocomms.PushNotificationsSendRequestToItem{
                SendPushNotificationDirectRecipient: &twiliocomms.SendPushNotificationDirectRecipient{
                    Token: "dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh",
                    Provider: twiliocomms.PushNotificationProviderFcm,
                },
            },
            &twiliocomms.PushNotificationsSendRequestToItem{
                SendPushNotificationDirectRecipient: &twiliocomms.SendPushNotificationDirectRecipient{
                    Token: "utJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFAukXD89hGq",
                    Provider: twiliocomms.PushNotificationProviderApn,
                },
            },
        },
        Content: &twiliocomms.PushNotificationContentBody{
            Title: twiliocomms.String(
                "Boarding time @ TLL",
            ),
            Body: "Dear customer, you have 1 hour until boarding time at the Tallinn airport",
        },
        Priority: twiliocomms.PushNotificationPriorityHigh.Ptr(),
        Sound: twiliocomms.String(
            "ring",
        ),
    }
    client.PushNotifications.Send(
        context.TODO(),
        request,
    )
}
