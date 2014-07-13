package geyser

import (
	"fmt"
)

var subscribeCache = make(map[string][]byte)
var unsubscribeCache = make(map[string][]byte)

var pingMsg = []byte(`{"event":"pusher:ping","data":{}}`)
var pongMsg = []byte(`{"event":"pusher:pong","data":{}}`)

func subscribe(channelName string) []byte {
	msg, ok := subscribeCache[channelName]
	if !ok {
		msg = []byte(fmt.Sprintf(`{"event":"pusher:subscribe", "data":"{\"channel\":\"%s\"}"}`, channelName))
		subscribeCache[channelName] = msg
	}
	return msg
}

func unsubscribe(channelName string) []byte {
	msg, ok := unsubscribeCache[channelName]
	if !ok {
		msg = []byte(fmt.Sprintf(`{"event":"pusher:unsubscribe", "data":"{\"channel\":\"%s\"}"}`, channelName))
		unsubscribeCache[channelName] = msg
	}
	return msg
}
