<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->credentials->fetch(
    'comms_credential_01h9krwprkeee8fzqspvwy6nq8',
);
