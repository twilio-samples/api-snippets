using TwilioComms;
using System.Threading.Tasks;
using System.Collections.Generic;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioCommsClient(
            accountId: "<username>",
            authToken: "<password>"
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
