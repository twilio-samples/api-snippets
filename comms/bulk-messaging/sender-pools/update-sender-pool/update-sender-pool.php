<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\SenderPools\Requests\SenderPoolsUpdateRequest;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
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
