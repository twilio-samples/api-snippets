from twilio import TwilioComms
import datetime

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.sender_pools.list(
    start_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    end_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    operation_id="comms_operation_01h9krwprkeee8fzqspvwy6nq8",
    page_token="pageToken",
    page_size=50,
)
