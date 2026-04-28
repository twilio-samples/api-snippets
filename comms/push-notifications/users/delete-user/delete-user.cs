using Twilio.Comms;
using System.Threading.Tasks;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.PushNotifications.Users.DeleteAsync(
            "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8"
        );
    }

}
