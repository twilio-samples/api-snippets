<?php

namespace Example;

use Twilio\Comms\TwilioComms;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->senderPools->removeSender(
    'comms_senderpool_01h9krwprkeee8fzqspvwy6nq8',
    'comms_sender_01h9krwprkeee8fzqspvwy6nq8',
);
