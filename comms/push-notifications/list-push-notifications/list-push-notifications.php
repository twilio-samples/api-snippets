<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\Requests\PushNotificationsListRequest;
use Twilio\Comms\Types\PushNotificationStatus;
use DateTime;
use Twilio\Comms\Types\PushNotificationProvider;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->list(
    new PushNotificationsListRequest([
        'operationId' => 'comms_operation_01h9krwprkeee8fzqspvwy6nq8',
        'status' => PushNotificationStatus::Scheduled->value,
        'tags' => 'key_1:value;key_2:value;',
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'user' => 'comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8',
        'provider' => PushNotificationProvider::Apn->value,
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
