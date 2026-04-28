from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.sender_pools.add_sender(
    sender_pool_id="comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
    request=[
        {
            "sender_id": "comms_sender_01h9krwprkeee8fzqspvwy6nq8"
        }
    ],
)
