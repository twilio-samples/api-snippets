<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\Senders\Requests\SendersSearchRequest;
use Twilio\Types\SenderCommunicationChannel;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->senders->search(
    new SendersSearchRequest([
        'address' => '+14153902337',
        'channel' => SenderCommunicationChannel::Sms->value,
    ]),
);
