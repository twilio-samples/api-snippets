require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.sender_pools.list_operations(
  start_date: "2024-01-15T09:30:00Z",
  end_date: "2024-01-15T09:30:00Z",
  status: "PROCESSING",
  page_token: "pageToken",
  page_size: 50
)
