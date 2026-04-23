<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\DeviceRegistrations\Requests\DeviceRegistrationsRegisterRequest;
use Twilio\Types\PushNotificationProvider;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->deviceRegistrations->register(
    new DeviceRegistrationsRegisterRequest([
        'userId' => 'comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8',
        'appName' => 'limonade_app',
        'token' => 'dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh',
        'provider' => PushNotificationProvider::Fcm->value,
    ]),
);
