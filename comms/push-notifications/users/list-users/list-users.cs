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

        await client.PushNotifications.Users.ListAsync(
            new UsersListRequest {
                PageToken = "pageToken",
                PageSize = 50
            }
        );
    }

}
