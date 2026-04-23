from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.messages.seek_by_identifier(
    identifier="identifier",
)
