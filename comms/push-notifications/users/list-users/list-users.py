from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.push_notifications.users.list(
    page_token="pageToken",
    page_size=50,
)
