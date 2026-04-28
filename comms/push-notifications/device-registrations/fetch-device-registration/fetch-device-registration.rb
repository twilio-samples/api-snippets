require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.push_notifications.device_registrations.fetch(device_registration_id: "comms_device_registration_01h9krwprkeee8fzqspvwy6nq8")
