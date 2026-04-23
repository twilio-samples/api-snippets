from twilio import TwilioComms, MessageContentTextWithMedia
from twilio.messages import MessagesSendRequestToItemAddress

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.messages.send(
    to=[
        MessagesSendRequestToItemAddress(
            address="+14153902337",
            channel="PHONE",
        )
    ],
    content=MessageContentTextWithMedia(
        text="Hello, World!",
    ),
)
