using Twilio.Comms;
using System.Threading.Tasks;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.SenderPools.RemoveSenderAsync(
            "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
            "comms_sender_01h9krwprkeee8fzqspvwy6nq8"
        );
    }

}
