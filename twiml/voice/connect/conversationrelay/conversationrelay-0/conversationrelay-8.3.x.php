<?php
require_once './vendor/autoload.php';
use Twilio\TwiML\VoiceResponse;

$response = new VoiceResponse();
$connect = $response->connect(['action' => 'https://myhttpserver.com/connect_action']);
$connect->conversation_relay(['url' => 'wss://mywebsocketserver.com/websocket', 'welcomeGreeting' => 'Hi! Ask me anything!']);

echo $response;
