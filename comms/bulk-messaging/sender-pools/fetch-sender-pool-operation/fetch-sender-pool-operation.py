from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.sender_pools.fetch_operation(
    operation_id="comms_operation_01h9krwprkeee8fzqspvwy6nq8",
)
