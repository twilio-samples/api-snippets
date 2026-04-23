using TwilioComms;
using System.Threading.Tasks;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioCommsClient(
            accountId: "<username>",
            authToken: "<password>"
        );

        await client.PushNotifications.Credentials.FetchAsync(
            "comms_credential_01h9krwprkeee8fzqspvwy6nq8"
        );
    }

}
