require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.push_notifications.list(
  operation_id: "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
  status: "SCHEDULED",
  tags: "key_1:value;key_2:value;",
  start_date: "2024-01-15T09:30:00Z",
  end_date: "2024-01-15T09:30:00Z",
  user: "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
  provider: "APN",
  page_token: "pageToken",
  page_size: 50
)
