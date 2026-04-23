<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->messages->fetch(
    'comms_message_01h9krwprkeee8fzqspvwy6nq8',
);
