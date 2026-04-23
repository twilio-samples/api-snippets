<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\Credentials\Requests\CredentialsCreateRequest;
use Twilio\Types\PushCredentialType;
use Twilio\Types\ApnCertificatePushCredential;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->credentials->create(
    new CredentialsCreateRequest([
        'credentialType' => PushCredentialType::Apn->value,
        'content' => new ApnCertificatePushCredential([
            'certificate' => 'LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tTUlJRm5UQ0NCSVdnQXdJQkFnSUlBank5SDg0OStFOHdEUVlKS29aSWh2Y05BUUVGQlFBd2daWXhDekFKQmdOVi4uLi4uQT09LS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ==',
            'privateKey' => 'LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLU1JSUVwUUlCQUFLQ0FRRUF1eWYvbE5ySDljazhEbU55bzNmR2d2Q0kxbDlzK2NtQlkzV0l6K2NVRHFteGlpZVIKLi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t',
        ]),
        'appName' => 'limonade_app',
    ]),
);
