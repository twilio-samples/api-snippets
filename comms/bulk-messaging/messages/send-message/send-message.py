from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.messages.send(
    to=[
        {
            "address": "+12065558844",
            "channel": "PHONE"
        }
    ],
    content={
        "text": "Hello, World!"
    },
    from_={
        "address": "+14153901002",
        "channel": "SMS"
    },
)
