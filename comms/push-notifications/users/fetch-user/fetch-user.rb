require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.push_notifications.users.fetch(user_id: "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8")
