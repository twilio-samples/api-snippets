<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\Senders\Requests\SendersListRequest;
use Twilio\Comms\Types\SenderCommunicationChannel;
use Twilio\Comms\Types\SenderStatus;
use DateTime;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
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
