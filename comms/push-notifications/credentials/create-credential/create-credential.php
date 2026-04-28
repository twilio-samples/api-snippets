<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\Credentials\Requests\CredentialsCreateRequest;
use Twilio\Comms\Types\PushCredentialType;
use Twilio\Comms\Types\ApnCertificatePushCredential;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
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
