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
