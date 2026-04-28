require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.messages.list(
  operation_id: "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
  session_id: "comms_session_01h9krwprkeee8fzqspvwy6nq8",
  start_date: "2024-01-15T09:30:00Z",
  end_date: "2024-01-15T09:30:00Z",
  profile: "mem_profile_01h9krwprkeee8fzqspvwy6nq8",
  channel: "SMS",
  status: "QUEUED",
  tags: "key_1:value;key_2:value;",
  page_token: "pageToken",
  page_size: 50
)
