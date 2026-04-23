require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.push_notifications.fetch_operation(operation_id: "comms_operation_01h9krwprkeee8fzqspvwy6nq8")
