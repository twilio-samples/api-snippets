<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\Requests\PushNotificationsListOperationsRequest;
use DateTime;
use Twilio\Comms\Types\CommunicationOperationStatus;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->listOperations(
    new PushNotificationsListOperationsRequest([
        'pageToken' => 'pageToken',
        'pageSize' => 50,
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'status' => CommunicationOperationStatus::Scheduled->value,
    ]),
);
