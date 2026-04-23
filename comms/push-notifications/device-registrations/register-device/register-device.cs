using TwilioComms;
using System.Threading.Tasks;
using TwilioComms.PushNotifications;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioCommsClient(
            accountId: "<username>",
            authToken: "<password>"
        );

        await client.PushNotifications.DeviceRegistrations.RegisterAsync(
            new DeviceRegistrationsRegisterRequest {
                UserId = "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
                AppName = "limonade_app",
                Token = "dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh",
                Provider = PushNotificationProvider.Fcm
            }
        );
    }

}
