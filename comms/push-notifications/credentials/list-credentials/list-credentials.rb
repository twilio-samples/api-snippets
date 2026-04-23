require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.push_notifications.credentials.list(
  start_date: "2024-01-15T09:30:00Z",
  end_date: "2024-01-15T09:30:00Z",
  app_name: "appName",
  credential_type: "APN",
  is_default: true,
  page_token: "pageToken",
  page_size: 50
)
