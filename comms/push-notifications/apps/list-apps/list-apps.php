<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\Apps\Requests\AppsListRequest;
use DateTime;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->apps->list(
    new AppsListRequest([
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'isDefault' => true,
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
