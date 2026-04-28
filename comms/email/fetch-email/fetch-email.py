from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.emails.fetch(
    email_id="comms_email_01h9krwprkeee8fzqspvwy6nq8",
)
