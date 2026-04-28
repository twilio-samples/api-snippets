<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\Users\Requests\UsersListRequest;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->users->list(
    new UsersListRequest([
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
