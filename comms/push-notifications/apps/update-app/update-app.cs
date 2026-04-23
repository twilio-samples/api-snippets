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

        await client.PushNotifications.Apps.UpdateAsync(
            "appName",
            new AppsUpdateRequest {
                IsDefault = true
            }
        );
    }

}
