<?php

namespace Example;

use Twilio\Comms\TwilioComms;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->fetchOperation(
    'comms_operation_01h9krwprkeee8fzqspvwy6nq8',
);
