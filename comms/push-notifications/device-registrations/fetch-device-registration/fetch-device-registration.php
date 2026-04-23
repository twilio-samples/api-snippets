<?php

namespace Example;

use Twilio\TwilioComms;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->deviceRegistrations->fetch(
    'comms_device_registration_01h9krwprkeee8fzqspvwy6nq8',
);
