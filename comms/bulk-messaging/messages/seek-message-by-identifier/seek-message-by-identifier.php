<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->messages->seekByIdentifier(
    'identifier',
);
