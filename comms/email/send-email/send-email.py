from twilio import TwilioComms, EmailAddressSender, EmailHtmlContent, EmailAttachmentsItem
from twilio.emails import EmailsSendRequestToItemAddress

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.emails.send(
    from_=EmailAddressSender(
        address="support@example.company.io",
        name="Cool Co Support",
    ),
    to=[
        EmailsSendRequestToItemAddress(
            address="bob@example.com",
            name="Bob Smith",
        )
    ],
    content=EmailHtmlContent(
        html="<html><body>Hey, <br/><br/>Cake</b></body></html>",
        text="Hey, the cake is ready.",
        subject="Re: Wedding Cake",
        attachments=[
            EmailAttachmentsItem(
                filename="filename",
                content_type="contentType",
                content="content",
            )
        ],
    ),
)
