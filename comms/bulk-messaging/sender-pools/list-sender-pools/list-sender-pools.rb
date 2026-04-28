require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.sender_pools.list(
  start_date: "2024-01-15T09:30:00Z",
  end_date: "2024-01-15T09:30:00Z",
  operation_id: "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
  page_token: "pageToken",
  page_size: 50
)
