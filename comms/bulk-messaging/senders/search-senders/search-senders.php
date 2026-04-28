<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\Senders\Requests\SendersSearchRequest;
use Twilio\Comms\Types\SenderCommunicationChannel;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->senders->search(
    new SendersSearchRequest([
        'address' => '+14153902337',
        'channel' => SenderCommunicationChannel::Sms->value,
    ]),
);
