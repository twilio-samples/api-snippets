<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\Messages\Requests\MessagesSendRequest;
use Twilio\Messages\Types\MessagesSendRequestToItemAddress;
use Twilio\Types\MessageAddressChannel;
use Twilio\Types\MessageContentTextWithMedia;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->messages->send(
    new MessagesSendRequest([
        'to' => [
            new MessagesSendRequestToItemAddress([
                'address' => '+14153902337',
                'channel' => MessageAddressChannel::Phone->value,
            ]),
        ],
        'content' => new MessageContentTextWithMedia([
            'text' => 'Hello, World!',
        ]),
    ]),
);
