using TwilioComms;
using System.Threading.Tasks;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioCommsClient(
            accountId: "<username>",
            authToken: "<password>"
        );

        await client.SenderPools.RemoveSenderAsync(
            "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
            "comms_sender_01h9krwprkeee8fzqspvwy6nq8"
        );
    }

}
