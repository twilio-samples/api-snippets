require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.senders.list(
  channel: "SMS",
  status: "ACTIVATED",
  start_date: "2024-01-15T09:30:00Z",
  end_date: "2024-01-15T09:30:00Z",
  page_token: "pageToken",
  page_size: 50
)
