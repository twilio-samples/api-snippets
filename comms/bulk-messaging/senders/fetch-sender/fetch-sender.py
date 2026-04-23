from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.senders.fetch(
    sender_id="comms_sender_01h9krwprkeee8fzqspvwy6nq8",
)
