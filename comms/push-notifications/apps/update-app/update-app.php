<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\Apps\Requests\AppsUpdateRequest;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->apps->update(
    'appName',
    new AppsUpdateRequest([
        'isDefault' => true,
    ]),
);
