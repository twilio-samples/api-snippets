<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\SenderPools\Requests\SenderPoolsListRequest;
use DateTime;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->senderPools->list(
    new SenderPoolsListRequest([
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'operationId' => 'comms_operation_01h9krwprkeee8fzqspvwy6nq8',
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
