<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\Senders\Types\SendersResolveRequestSenderPoolId;
use Twilio\Types\CommunicationRecipient;
use Twilio\Types\CommunicationChannel;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->senders->resolve(
    new SendersResolveRequestSenderPoolId([
        'recipientAddresses' => [
            new CommunicationRecipient([
                'address' => '+14153902337',
                'channel' => CommunicationChannel::Phone->value,
            ]),
            new CommunicationRecipient([
                'address' => '+14153902337',
                'channel' => CommunicationChannel::Whatsapp->value,
            ]),
            new CommunicationRecipient([
                'address' => 'davidpletnjov@example.com',
                'channel' => CommunicationChannel::Email->value,
            ]),
        ],
    ]),
);
