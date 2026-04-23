<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->emails->delete(
    'comms_email_01h9krwprkeee8fzqspvwy6nq8',
);
