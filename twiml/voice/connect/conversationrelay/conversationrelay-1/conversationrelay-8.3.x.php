<?php
require_once './vendor/autoload.php';
use Twilio\TwiML\VoiceResponse;

$response = new VoiceResponse();
$connect = $response->connect();
$conversationrelay = $connect->conversation_relay(['url' => 'wss://mywebsocketserver.com/websocket']);
$conversationrelay->language(['code' => 'sv-SE', 'ttsProvider' => 'amazon', 'voice' => 'Elin-Neural', 'transcriptionProvider' => 'google', 'speechModel' => 'long']);
$conversationrelay->language(['code' => 'en-US', 'ttsProvider' => 'google', 'voice' => 'en-US-Journey-O']);

echo $response;
