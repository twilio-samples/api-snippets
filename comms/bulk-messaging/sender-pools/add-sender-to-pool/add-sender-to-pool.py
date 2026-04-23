from twilio import TwilioComms
from twilio.sender_pools import SenderPoolsAddSenderRequestItem

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.sender_pools.add_sender(
    sender_pool_id="comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
    request=[
        SenderPoolsAddSenderRequestItem(
            sender_id="comms_sender_01h9krwprkeee8fzqspvwy6nq8",
        )
    ],
)
