using Twilio.Comms;
using System.Threading.Tasks;
using Twilio.Comms.PushNotifications;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.PushNotifications.Users.ListAsync(
            new UsersListRequest {
                PageToken = "pageToken",
                PageSize = 50
            }
        );
    }

}
