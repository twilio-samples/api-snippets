<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\Credentials\Requests\CredentialsUpdateRequest;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->credentials->update(
    'comms_credential_01h9krwprkeee8fzqspvwy6nq8',
    new CredentialsUpdateRequest([
        'isDefault' => true,
    ]),
);
