require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.push_notifications.device_registrations.fetch(device_registration_id: "comms_device_registration_01h9krwprkeee8fzqspvwy6nq8")
