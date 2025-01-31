<?php
require_once './vendor/autoload.php';
use Twilio\TwiML\VoiceResponse;

$response = new VoiceResponse();
$connect = $response->connect();
$conversationrelay = $connect->conversation_relay(['url' => 'wss://mywebsocketserver.com/websocket']);
$conversationrelay->parameter(['name' => 'foo', 'value' => 'bar']);
$conversationrelay->parameter(['name' => 'hint', 'value' => 'Annoyed customer']);

echo $response;
