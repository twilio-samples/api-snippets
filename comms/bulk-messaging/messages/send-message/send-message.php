<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\Messages\Requests\MessagesSendRequest;
use Twilio\Comms\Messages\Types\MessagesSendRequestToItemAddress;
use Twilio\Comms\Types\MessageAddressChannel;
use Twilio\Comms\Types\MessageContentTextWithMedia;
use Twilio\Comms\Types\MessageAddressSender;
use Twilio\Comms\Types\MessageSenderChannel;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->messages->send(
    new MessagesSendRequest([
        'to' => [
            new MessagesSendRequestToItemAddress([
                'address' => '+12065558844',
                'channel' => MessageAddressChannel::Phone->value,
            ]),
        ],
        'content' => new MessageContentTextWithMedia([
            'text' => 'Hello, World!',
        ]),
        'from' => new MessageAddressSender([
            'address' => '+14153901002',
            'channel' => MessageSenderChannel::Sms->value,
        ]),
    ]),
);
