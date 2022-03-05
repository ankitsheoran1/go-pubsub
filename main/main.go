package main

import (
"fmt"
	pubsub "go-pubsub-master"
)

func main() {
	ps := pubsub.New()
	subscriber := ps.Subscribe("user:1", "user:2").Do(func(message pubsub.Message) {
		println("We got meessage on topic", message.Topic())
	})

	defer subscriber.Close()

	count := ps.Publish(pubsub.NewMessage("user:1:username", "hello"))

	fmt.Printf("Published messages to %v subscribers", count)
}

