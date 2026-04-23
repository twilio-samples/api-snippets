require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.messages.fetch(message_id: "comms_message_01h9krwprkeee8fzqspvwy6nq8")
