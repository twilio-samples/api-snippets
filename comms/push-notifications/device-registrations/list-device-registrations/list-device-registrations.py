from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.push_notifications.device_registrations.list(
    user_id="comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
    app_name="appName",
    page_token="pageToken",
    page_size=50,
)
