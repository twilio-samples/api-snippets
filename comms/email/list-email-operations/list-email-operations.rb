require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.emails.list_operations(
  start_date: "2024-01-15T09:30:00Z",
  end_date: "2024-01-15T09:30:00Z",
  status: "SCHEDULED",
  page_token: "pageToken",
  page_size: 50
)
