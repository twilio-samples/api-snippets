require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.push_notifications.fetch(push_notification_id: "comms_pushnotification_01h9krwprkeee8fzqspvwy6nq8")
