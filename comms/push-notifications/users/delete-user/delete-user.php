<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->users->delete(
    'comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8',
);
