<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\SenderPools\Requests\SenderPoolsUpdateRequest;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->senderPools->update(
    'comms_senderpool_01h9krwprkeee8fzqspvwy6nq8',
    new SenderPoolsUpdateRequest([
        'name' => 'Customer Support Senders - Tier 1',
        'tags' => [
            'tier' => '1',
        ],
    ]),
);
