from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.push_notifications.users.delete(
    user_id="comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
)
