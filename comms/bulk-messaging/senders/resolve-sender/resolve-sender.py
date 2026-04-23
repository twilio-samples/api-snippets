from twilio import TwilioComms, CommunicationRecipient
from twilio.senders import SendersResolveRequestSenderPoolId

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.senders.resolve(
    request=SendersResolveRequestSenderPoolId(
        recipient_addresses=[
            CommunicationRecipient(
                address="+14153902337",
                channel="PHONE",
            ),
            CommunicationRecipient(
                address="+14153902337",
                channel="WHATSAPP",
            ),
            CommunicationRecipient(
                address="davidpletnjov@example.com",
                channel="EMAIL",
            )
        ],
    ),
)
