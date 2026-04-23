require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.sender_pools.list_senders(
  sender_pool_id: "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
  channel: "SMS",
  status: "ACTIVATED",
  start_date: "2024-01-15T09:30:00Z",
  end_date: "2024-01-15T09:30:00Z",
  page_token: "pageToken",
  page_size: 50
)
