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

        await client.PushNotifications.Credentials.UpdateAsync(
            "comms_credential_01h9krwprkeee8fzqspvwy6nq8",
            new CredentialsUpdateRequest {
                IsDefault = true
            }
        );
    }

}
