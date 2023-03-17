package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

var rdb3 = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       3,
})

var ctx = context.Background()

func Receive(groupName, sender string, signal <-chan bool) {
	pubsub := rdb3.Subscribe(ctx, groupName)
	defer pubsub.Close()
	select {
	case <-signal:
		return
	default:

		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}
			if !strings.Contains(msg.Payload, sender) {
				fmt.Println(strings.Split(msg.Payload, " ")[0] + " : " + strings.Split(msg.Payload, " ")[1])
			}
		}
	}
}
