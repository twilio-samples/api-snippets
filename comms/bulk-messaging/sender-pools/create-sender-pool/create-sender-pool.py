from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.sender_pools.create(
    name="Sales Leads - APAC",
    tags={
        "region": "APAC"
    },
)
