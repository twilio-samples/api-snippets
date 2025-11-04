<?php
require_once './vendor/autoload.php';
use Twilio\TwiML\VoiceResponse;

$response = new VoiceResponse();
$start = $response->start();
$start->recording(['channels' => 'dual',
    'recordingStatusCallback' => 'https://example.com/your-callback-url']);
$response->say('The recording has started.');

echo $response;
