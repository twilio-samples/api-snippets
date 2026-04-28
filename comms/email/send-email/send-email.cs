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

        await client.Emails.SendAsync(
            new EmailsSendRequest {
                From = new EmailAddressSender {
                    Address = "support@example.company.io",
                    Name = "Cool Co Support"
                },
                To = new List<OneOf<EmailsSendRequestToItemAddress, EmailsSendRequestToItemMemoryStoreId>>(){
                    new EmailsSendRequestToItemAddress {
                        Address = "bob@example.com",
                        Name = "Bob Smith"
                    },
                }
                ,
                Content = new EmailHtmlContent {
                    Html = "<html><body>Hey, <br/><br/><b>Cake</b></body></html>",
                    Text = "Hey, the cake is ready.",
                    Subject = "Re: Wedding Cake",
                    Attachments = new List<EmailAttachmentsItem>(){
                        new EmailAttachmentsItem {
                            Filename = "filename",
                            ContentType = "contentType",
                            Content = "content"
                        },
                    }

                }
            }
        );
    }

}
