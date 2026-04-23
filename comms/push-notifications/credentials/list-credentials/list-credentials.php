<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\Credentials\Requests\CredentialsListRequest;
use DateTime;
use Twilio\Types\PushCredentialType;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
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
