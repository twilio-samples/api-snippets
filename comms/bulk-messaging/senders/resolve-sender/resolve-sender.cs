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

        await client.Senders.ResolveAsync(
            new SendersResolveRequestSenderPoolId {
                RecipientAddresses = new List<CommunicationRecipient>(){
                    new CommunicationRecipient {
                        Address = "+14153902337",
                        Channel = CommunicationChannel.Phone
                    },
                    new CommunicationRecipient {
                        Address = "+14153902337",
                        Channel = CommunicationChannel.Whatsapp
                    },
                    new CommunicationRecipient {
                        Address = "davidpletnjov@example.com",
                        Channel = CommunicationChannel.Email
                    },
                }

            }
        );
    }

}
