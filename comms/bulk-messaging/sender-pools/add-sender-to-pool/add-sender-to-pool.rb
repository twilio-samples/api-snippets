require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.sender_pools.add_sender(
  sender_pool_id: "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
  request: [{
    sender_id: "comms_sender_01h9krwprkeee8fzqspvwy6nq8"
  }]
)
