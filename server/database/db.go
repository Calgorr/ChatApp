package database

import (
	"context"
	"errors"
	"time"

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

var rdb1 = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       1,
})

var rdb2 = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       2,
})
var rdb3 = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       2,
})

func AddUser(user *model.User) error {
	if rdb0.Get(ctx, user.Username).Val() != "" {
		return errors.New("user already exists")
	}
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

func AddMessage(message *model.Message) error {
	return rdb1.Set(ctx, message.GroupName+time.Now().String(), message, 1*time.Hour).Err()
}

func AddGroup(group *model.Group) error {
	return rdb2.Set(ctx, group.GroupName, group, 0).Err()
}

func GetGroups(username string) ([]model.Group, error) {
	cmd := rdb2.Do(ctx, "keys", "*")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	var groups []model.Group
	for _, key := range cmd.Val().([]interface{}) {
		group := &model.Group{}
		rdb2.Get(ctx, key.(string)).Scan(group)
		if !group.CheckMember(*&model.User{Username: username}) {
			groups = append(groups, *group)
		}
	}
	return groups, nil
}

func AddMemberToGroup(groupName string, user *model.User) error {
	group := &model.Group{}
	rdb2.Get(ctx, groupName).Scan(group)
	if group.CheckMember(*user) {
		return errors.New("user already exists in group")
	}
	group.AddMember(*user)
	return rdb2.Set(ctx, groupName, group, 0).Err()
}

func Publish(message *model.Message) error {
	return rdb3.Publish(ctx, message.GroupName, message.Sender+" "+message.Content).Err()
}

func GetMessages(groupName string) ([]model.Message, error) {
	cmd := rdb1.Do(ctx, "keys", groupName+"*")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	var messages []model.Message
	for _, key := range cmd.Val().([]interface{}) {
		message := &model.Message{}
		rdb1.Get(ctx, key.(string)).Scan(message)
		if message.GroupName == groupName {
			messages = append(messages, *message)
		}
	}
	return messages, nil
}

func CheckGroup(groupName string) bool {
	return rdb2.Get(ctx, groupName).Val() != ""
}

func GetGroup(groupName string) (*model.Group, error) {
	group := &model.Group{}
	rdb2.Get(ctx, groupName).Scan(group)
	return group, nil
}
