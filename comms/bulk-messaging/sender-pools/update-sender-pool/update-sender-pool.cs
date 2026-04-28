using Twilio.Comms;
using System.Threading.Tasks;
using System.Collections.Generic;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.SenderPools.UpdateAsync(
            "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
            new SenderPoolsUpdateRequest {
                Name = "Customer Support Senders - Tier 1",
                Tags = new Dictionary<string, string>(){
                    ["tier"] = "1",
                }

            }
        );
    }

}
