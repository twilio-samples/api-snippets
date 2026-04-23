<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\Messages\Requests\MessagesListOperationsRequest;
use DateTime;
use Twilio\Types\CommunicationOperationStatus;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->messages->listOperations(
    new MessagesListOperationsRequest([
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'status' => CommunicationOperationStatus::Scheduled->value,
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
