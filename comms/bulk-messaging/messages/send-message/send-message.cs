using TwilioComms;
using System.Threading.Tasks;
using System.Collections.Generic;
using OneOf;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioCommsClient(
            accountId: "<username>",
            authToken: "<password>"
        );

        await client.Messages.SendAsync(
            new MessagesSendRequest {
                To = new List<OneOf<MessagesSendRequestToItemAddress, MessagesSendRequestToItemAddresses, MessagesSendRequestToItemMemoryStoreId>>(){
                    new MessagesSendRequestToItemAddress {
                        Address = "+14153902337",
                        Channel = MessageAddressChannel.Phone
                    },
                }
                ,
                Content = new MessageContentTextWithMedia {
                    Text = "Hello, World!"
                }
            }
        );
    }

}
