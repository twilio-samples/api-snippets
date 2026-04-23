<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->fetchOperation(
    'comms_operation_01h9krwprkeee8fzqspvwy6nq8',
);
