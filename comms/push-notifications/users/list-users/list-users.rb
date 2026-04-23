require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.push_notifications.users.list(
  page_token: "pageToken",
  page_size: 50
)
