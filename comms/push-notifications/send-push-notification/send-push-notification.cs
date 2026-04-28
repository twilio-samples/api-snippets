using Twilio.Comms;
using System.Threading.Tasks;
using System.Collections.Generic;
using OneOf;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.PushNotifications.SendAsync(
            new PushNotificationsSendRequest {
                From = new PushNotificationCredentialSender {
                    Fcm = "comms_credential_01h9krwprkeee8fzqspvwy6nq8",
                    Apn = "comms_credential_8qn6ywvpsqzf8eeekrpwrk9h10"
                },
                To = new List<OneOf<SendPushNotificationDirectRecipient, SendPushNotificationsUserRecipient>>(){
                    new SendPushNotificationDirectRecipient {
                        Token = "dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh",
                        Provider = PushNotificationProvider.Fcm
                    },
                    new SendPushNotificationDirectRecipient {
                        Token = "utJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFAukXD89hGq",
                        Provider = PushNotificationProvider.Apn
                    },
                }
                ,
                Content = new PushNotificationContentBody {
                    Title = "Boarding time @ TLL",
                    Body = "Dear customer, you have 1 hour until boarding time at the Tallinn airport"
                },
                Priority = PushNotificationPriority.High,
                Sound = "ring"
            }
        );
    }

}
