<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\DeviceRegistrations\Requests\DeviceRegistrationsListRequest;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->deviceRegistrations->list(
    new DeviceRegistrationsListRequest([
        'userId' => 'comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8',
        'appName' => 'appName',
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
