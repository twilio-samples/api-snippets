<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\SenderPools\Types\SenderPoolsAddSenderRequestItem;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->senderPools->addSender(
    'comms_senderpool_01h9krwprkeee8fzqspvwy6nq8',
    [
        new SenderPoolsAddSenderRequestItem([
            'senderId' => 'comms_sender_01h9krwprkeee8fzqspvwy6nq8',
        ]),
    ],
);
