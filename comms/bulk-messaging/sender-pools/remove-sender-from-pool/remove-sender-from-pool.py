from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.sender_pools.remove_sender(
    sender_pool_id="comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
    sender_id="comms_sender_01h9krwprkeee8fzqspvwy6nq8",
)
