<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\Apps\Requests\AppsUpdateRequest;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->apps->update(
    'appName',
    new AppsUpdateRequest([
        'isDefault' => true,
    ]),
);
