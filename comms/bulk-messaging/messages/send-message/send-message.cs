using Twilio.Comms;
using System.Threading.Tasks;
using System.Collections.Generic;
using OneOf;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.Messages.SendAsync(
            new MessagesSendRequest {
                To = new List<OneOf<MessagesSendRequestToItemAddress, MessagesSendRequestToItemAddresses, MessagesSendRequestToItemMemoryStoreId>>(){
                    new MessagesSendRequestToItemAddress {
                        Address = "+12065558844",
                        Channel = MessageAddressChannel.Phone
                    },
                }
                ,
                Content = new MessageContentTextWithMedia {
                    Text = "Hello, World!"
                },
                From = new MessageAddressSender {
                    Address = "+14153901002",
                    Channel = MessageSenderChannel.Sms
                }
            }
        );
    }

}
