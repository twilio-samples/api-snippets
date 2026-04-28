from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.push_notifications.credentials.fetch(
    credential_id="comms_credential_01h9krwprkeee8fzqspvwy6nq8",
)
