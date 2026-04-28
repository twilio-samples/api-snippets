from twilio_comms import TwilioComms
import datetime

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.sender_pools.list_senders(
    sender_pool_id="comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
    channel="SMS",
    status="ACTIVATED",
    start_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    end_date=datetime.datetime.fromisoformat("2024-01-15T09:30:00+00:00"),
    page_token="pageToken",
    page_size=50,
)
