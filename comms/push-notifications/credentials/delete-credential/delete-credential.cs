using TwilioComms;
using System.Threading.Tasks;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioCommsClient(
            accountId: "<username>",
            authToken: "<password>"
        );

        await client.PushNotifications.Credentials.DeleteAsync(
            "comms_credential_01h9krwprkeee8fzqspvwy6nq8"
        );
    }

}
