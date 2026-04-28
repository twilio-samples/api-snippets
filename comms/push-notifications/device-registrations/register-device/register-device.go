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
    request := &twiliocomms.DeviceRegistrationsRegisterRequest{
        UserId: twiliocomms.String(
            "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
        ),
        AppName: "limonade_app",
        Token: "dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh",
        Provider: twiliocomms.PushNotificationProviderFcm,
    }
    client.PushNotifications.DeviceRegistrations.Register(
        context.TODO(),
        request,
    )
}
