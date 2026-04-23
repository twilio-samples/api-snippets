from twilio import TwilioComms
import datetime

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.push_notifications.apps.list(
    start_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    end_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    is_default=True,
    page_token="pageToken",
    page_size=50,
)
