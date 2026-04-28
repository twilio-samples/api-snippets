package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.pushnotifications.requests.PushNotificationsSendRequest;
import com.twilio.sdk.resources.pushnotifications.types.PushNotificationsSendRequestFrom;
import com.twilio.sdk.resources.pushnotifications.types.PushNotificationsSendRequestToItem;
import com.twilio.sdk.types.PushNotificationContentBody;
import com.twilio.sdk.types.PushNotificationCredentialSender;
import com.twilio.sdk.types.PushNotificationPriority;
import com.twilio.sdk.types.PushNotificationProvider;
import com.twilio.sdk.types.SendPushNotificationDirectRecipient;
import java.util.Arrays;
import java.util.Optional;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.pushNotifications().send(
            PushNotificationsSendRequest
                .builder()
                .content(
                    PushNotificationContentBody
                        .builder()
                        .body("Dear customer, you have 1 hour until boarding time at the Tallinn airport")
                        .title("Boarding time @ TLL")
                        .build()
                )
                .to(
                    Arrays.asList(
                        PushNotificationsSendRequestToItem.of(
                            SendPushNotificationDirectRecipient
                                .builder()
                                .token("dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh")
                                .provider(PushNotificationProvider.FCM)
                                .build()
                        ),
                        PushNotificationsSendRequestToItem.of(
                            SendPushNotificationDirectRecipient
                                .builder()
                                .token("utJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFAukXD89hGq")
                                .provider(PushNotificationProvider.APN)
                                .build()
                        )
                    )
                )
                .from(
                    PushNotificationsSendRequestFrom.of(
                        PushNotificationCredentialSender
                            .builder()
                            .fcm(Optional.of("comms_credential_01h9krwprkeee8fzqspvwy6nq8"))
                            .apn(Optional.of("comms_credential_8qn6ywvpsqzf8eeekrpwrk9h10"))
                            .build()
                    )
                )
                .priority(PushNotificationPriority.HIGH)
                .sound("ring")
                .build()
        );
    }
}