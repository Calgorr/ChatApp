package handler

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var rdb3 = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       3,
})

var ctx = context.Background()

func Recive(groupName string) {
	pubsub := rdb3.Subscribe(ctx, groupName)
	defer pubsub.Close()
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		println(msg.Payload)
	}
}
