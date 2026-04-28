<?php

namespace Example;

use Twilio\Comms\TwilioComms;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->senderPools->delete(
    'comms_senderpool_01h9krwprkeee8fzqspvwy6nq8',
);
