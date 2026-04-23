from twilio import TwilioComms

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.push_notifications.device_registrations.delete(
    device_registration_id="comms_device_registration_01h9krwprkeee8fzqspvwy6nq8",
)
