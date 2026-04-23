<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->senders->fetch(
    'comms_sender_01h9krwprkeee8fzqspvwy6nq8',
);
