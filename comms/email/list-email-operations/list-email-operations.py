from twilio import TwilioComms
import datetime

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.emails.list_operations(
    start_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    end_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    status="SCHEDULED",
    page_token="pageToken",
    page_size=50,
)
