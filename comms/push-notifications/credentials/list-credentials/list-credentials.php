<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\Credentials\Requests\CredentialsListRequest;
use DateTime;
use Twilio\Comms\Types\PushCredentialType;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->credentials->list(
    new CredentialsListRequest([
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'appName' => 'appName',
        'credentialType' => PushCredentialType::Apn->value,
        'isDefault' => true,
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
