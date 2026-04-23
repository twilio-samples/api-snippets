from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.sender_pools.fetch(
    sender_pool_id="comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
)
