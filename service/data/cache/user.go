package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	redis8 "github.com/gomodule/redigo/redis"
	"thor/service/data"
)

const (
	USER_LOGIN_NAME_KEY = "blade:user:info:name:%s"
	TTL                 = 24 * 60 * 60
)

type User struct {
	Uid      uint64 `json:"uid"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Sn       string `json:"sn"`
	Identity uint64 `json:"identity"`
	GiveName string `json:"give_name"`
	Token    string `json:"token"`
}

type UserCache struct {
	rdb redis8.Conn
}

// func NewUserModel(logger *log.Logger) *userModel {
func NewUserCache() *UserCache {
	return &UserCache{
		rdb: data.GetBladeRedis(),
	}
}

func (t *UserCache) getUserKeyByName(username string) string {
	key := fmt.Sprintf(USER_LOGIN_NAME_KEY, username)
	return key
}

func (t *UserCache) GetUserInfoByName(username string) (*User, error) {
	if username == "" {
		return nil, errors.New("params errors")
	}

	result, err := redis8.String(t.rdb.Do("GET", t.getUserKeyByName(username)))
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal([]byte(result), user); err != nil {
		return nil, err
	}
	return user, nil
}

func (t *UserCache) SetUserInfoByName(user *User) (bool, error) {
	if user.UserName == "" {
		return false, errors.New("params errors!")
	}

	userStr, err := json.Marshal(user)
	if err != nil {
		return false, err
	}

	result, err := redis8.String(t.rdb.Do("SET", t.getUserKeyByName(user.UserName), userStr))
	if err != nil {
		return false, err
	}

	if result == "OK" {
		return true, nil
	}

	return false, errors.New("set errors")
}
