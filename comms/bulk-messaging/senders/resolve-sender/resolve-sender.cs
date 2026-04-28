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
