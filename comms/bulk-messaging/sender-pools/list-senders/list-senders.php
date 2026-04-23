<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\SenderPools\Requests\SenderPoolsListSendersRequest;
use Twilio\Types\SenderCommunicationChannel;
use Twilio\Types\SenderStatus;
use DateTime;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->senderPools->listSenders(
    'comms_senderpool_01h9krwprkeee8fzqspvwy6nq8',
    new SenderPoolsListSendersRequest([
        'channel' => SenderCommunicationChannel::Sms->value,
        'status' => SenderStatus::Activated->value,
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
