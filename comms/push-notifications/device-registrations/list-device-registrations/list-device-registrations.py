from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.push_notifications.device_registrations.list(
    user_id="comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
    app_name="appName",
    page_token="pageToken",
    page_size=50,
)
