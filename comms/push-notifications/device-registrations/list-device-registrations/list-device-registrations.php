<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\DeviceRegistrations\Requests\DeviceRegistrationsListRequest;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->deviceRegistrations->list(
    new DeviceRegistrationsListRequest([
        'userId' => 'comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8',
        'appName' => 'appName',
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
