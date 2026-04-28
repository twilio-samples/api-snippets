from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.push_notifications.fetch(
    push_notification_id="comms_pushnotification_01h9krwprkeee8fzqspvwy6nq8",
)
