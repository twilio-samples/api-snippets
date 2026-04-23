<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\Senders\Requests\SendersListRequest;
use Twilio\Types\SenderCommunicationChannel;
use Twilio\Types\SenderStatus;
use DateTime;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->senders->list(
    new SendersListRequest([
        'channel' => SenderCommunicationChannel::Sms->value,
        'status' => SenderStatus::Activated->value,
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
