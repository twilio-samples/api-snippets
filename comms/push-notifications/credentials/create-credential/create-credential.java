package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.pushnotifications.credentials.requests.CredentialsCreateRequest;
import com.twilio.sdk.types.ApnCertificatePushCredential;
import com.twilio.sdk.types.PushCredentialContent;
import com.twilio.sdk.types.PushCredentialType;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.pushNotifications().credentials().create(
            CredentialsCreateRequest
                .builder()
                .credentialType(PushCredentialType.APN)
                .content(
                    PushCredentialContent.of(
                        ApnCertificatePushCredential
                            .builder()
                            .certificate("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tTUlJRm5UQ0NCSVdnQXdJQkFnSUlBank5SDg0OStFOHdEUVlKS29aSWh2Y05BUUVGQlFBd2daWXhDekFKQmdOVi4uLi4uQT09LS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ==")
                            .privateKey("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLU1JSUVwUUlCQUFLQ0FRRUF1eWYvbE5ySDljazhEbU55bzNmR2d2Q0kxbDlzK2NtQlkzV0l6K2NVRHFteGlpZVIKLi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t")
                            .build()
                    )
                )
                .appName("limonade_app")
                .build()
        );
    }
}