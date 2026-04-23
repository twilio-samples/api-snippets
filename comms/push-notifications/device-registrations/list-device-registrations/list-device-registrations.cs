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

        await client.PushNotifications.DeviceRegistrations.ListAsync(
            new DeviceRegistrationsListRequest {
                UserId = "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
                AppName = "appName",
                PageToken = "pageToken",
                PageSize = 50
            }
        );
    }

}
