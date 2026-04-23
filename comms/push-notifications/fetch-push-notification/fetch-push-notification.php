<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->fetch(
    'comms_pushnotification_01h9krwprkeee8fzqspvwy6nq8',
);
