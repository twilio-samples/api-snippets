<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\SenderPools\Requests\SenderPoolsCreateRequest;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->senderPools->create(
    new SenderPoolsCreateRequest([
        'name' => 'Sales Leads - APAC',
        'tags' => [
            'region' => 'APAC',
        ],
    ]),
);
