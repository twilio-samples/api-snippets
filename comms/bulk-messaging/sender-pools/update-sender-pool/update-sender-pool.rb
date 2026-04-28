require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.sender_pools.update(
  sender_pool_id: "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
  name: "Customer Support Senders - Tier 1",
  tags: {
    tier: "1"
  }
)
