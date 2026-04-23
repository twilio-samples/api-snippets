from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.messages.fetch(
    message_id="comms_message_01h9krwprkeee8fzqspvwy6nq8",
)
