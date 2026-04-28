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

        await client.SenderPools.CreateAsync(
            new SenderPoolsCreateRequest {
                Name = "Sales Leads - APAC",
                Tags = new Dictionary<string, string>(){
                    ["region"] = "APAC",
                }

            }
        );
    }

}
