<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\Apps\Requests\AppsListRequest;
use DateTime;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
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
