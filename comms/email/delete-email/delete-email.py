from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.emails.delete(
    email_id="comms_email_01h9krwprkeee8fzqspvwy6nq8",
)
