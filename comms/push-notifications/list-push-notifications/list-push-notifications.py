from twilio import TwilioComms
import datetime

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.push_notifications.list(
    operation_id="comms_operation_01h9krwprkeee8fzqspvwy6nq8",
    status="SCHEDULED",
    tags="key_1:value;key_2:value;",
    start_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    end_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    user="comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
    provider="APN",
    page_token="pageToken",
    page_size=50,
)
