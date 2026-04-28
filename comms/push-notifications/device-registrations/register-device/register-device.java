package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.pushnotifications.deviceregistrations.requests.DeviceRegistrationsRegisterRequest;
import com.twilio.sdk.types.PushNotificationProvider;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.pushNotifications().deviceRegistrations().register(
            DeviceRegistrationsRegisterRequest
                .builder()
                .appName("limonade_app")
                .token("dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh")
                .provider(PushNotificationProvider.FCM)
                .userId("comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8")
                .build()
        );
    }
}