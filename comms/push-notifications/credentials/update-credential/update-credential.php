<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\Credentials\Requests\CredentialsUpdateRequest;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->credentials->update(
    'comms_credential_01h9krwprkeee8fzqspvwy6nq8',
    new CredentialsUpdateRequest([
        'isDefault' => true,
    ]),
);
