package database

import (
	"context"
	"errors"

	"github.com/Calgorr/ChatApp/server/model"
	redis "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// user database
var rdb0 = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func AddUser(user *model.User) error {
	return rdb0.Set(ctx, user.Username, user.Password, 0).Err()
}

func GetUser(username, password string) (*model.User, error) {
	val, err := rdb0.Get(ctx, username).Result()
	if err != nil {
		return nil, err
	}
	if val != password {
		return nil, errors.New("password is incorrect")
	}
	user := &model.User{
		Username: username,
		Password: val,
	}
	return user, nil
}
