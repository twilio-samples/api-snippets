<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\SenderPools\Requests\SenderPoolsListOperationsRequest;
use DateTime;
use Twilio\Comms\Types\OperationStatus;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->senderPools->listOperations(
    new SenderPoolsListOperationsRequest([
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'status' => OperationStatus::Processing->value,
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
