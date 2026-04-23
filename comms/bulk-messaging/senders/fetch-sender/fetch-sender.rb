require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.senders.fetch(sender_id: "comms_sender_01h9krwprkeee8fzqspvwy6nq8")
