<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->senderPools->removeSender(
    'comms_senderpool_01h9krwprkeee8fzqspvwy6nq8',
    'comms_sender_01h9krwprkeee8fzqspvwy6nq8',
);
