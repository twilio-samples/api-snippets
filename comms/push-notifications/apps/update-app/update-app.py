from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.push_notifications.apps.update(
    app_name="appName",
    is_default=True,
)
