<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\SenderPools\Requests\SenderPoolsListOperationsRequest;
use DateTime;
use Twilio\Types\OperationStatus;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
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
