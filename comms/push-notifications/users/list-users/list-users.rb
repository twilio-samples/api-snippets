require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.push_notifications.users.list(
  page_token: "pageToken",
  page_size: 50
)
