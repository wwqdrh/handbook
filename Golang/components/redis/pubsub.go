package redis

import (
	"fmt"
)

func TryPub(channel, message string) {
	err := client.Publish(ctx, channel, message).Err()
	if err != nil {
		fmt.Println("发生错误")
	}
}

func TrySub(channel string) string {
	sub := client.Subscribe(ctx, channel)
	defer sub.Close()

	// for {
	// 	msg, err := sub.ReceiveMessage(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(msg.Channel, msg.Payload)
	// }
	ch := sub.Channel()
	var message string
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		message = msg.Payload
	}
	return message
}
