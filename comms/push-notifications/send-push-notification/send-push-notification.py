from twilio import TwilioComms, PushNotificationCredentialSender, SendPushNotificationDirectRecipient, PushNotificationContentBody

client = TwilioComms(
    account_id="<username>",
    auth_token="<password>",
)

client.push_notifications.send(
    from_=PushNotificationCredentialSender(
        fcm="comms_credential_01h9krwprkeee8fzqspvwy6nq8",
        apn="comms_credential_8qn6ywvpsqzf8eeekrpwrk9h10",
    ),
    to=[
        SendPushNotificationDirectRecipient(
            token="dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh",
            provider="FCM",
        ),
        SendPushNotificationDirectRecipient(
            token="utJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFAukXD89hGq",
            provider="APN",
        )
    ],
    content=PushNotificationContentBody(
        title="Boarding time @ TLL",
        body="Dear customer, you have 1 hour until boarding time at the Tallinn airport",
    ),
    priority="HIGH",
    sound="ring",
)
