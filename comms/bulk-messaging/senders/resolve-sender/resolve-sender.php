<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\Senders\Types\SendersResolveRequestSenderPoolId;
use Twilio\Comms\Types\CommunicationRecipient;
use Twilio\Comms\Types\CommunicationChannel;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
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
