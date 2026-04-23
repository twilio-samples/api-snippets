from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.push_notifications.credentials.update(
    credential_id="comms_credential_01h9krwprkeee8fzqspvwy6nq8",
    is_default=True,
)
