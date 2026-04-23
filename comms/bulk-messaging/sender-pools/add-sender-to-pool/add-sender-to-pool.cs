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

        await client.SenderPools.AddSenderAsync(
            "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
            new List<SenderPoolsAddSenderRequestItem>(){
                new SenderPoolsAddSenderRequestItem {
                    SenderId = "comms_sender_01h9krwprkeee8fzqspvwy6nq8"
                },
            }
        );
    }

}
