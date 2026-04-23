using TwilioComms;
using System.Threading.Tasks;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioCommsClient(
            accountId: "<username>",
            authToken: "<password>"
        );

        await client.Senders.SearchAsync(
            new SendersSearchRequest {
                Address = "+14153902337",
                Channel = SenderCommunicationChannel.Sms
            }
        );
    }

}
