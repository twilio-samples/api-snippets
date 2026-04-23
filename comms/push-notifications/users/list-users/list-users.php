<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\Users\Requests\UsersListRequest;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->users->list(
    new UsersListRequest([
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
