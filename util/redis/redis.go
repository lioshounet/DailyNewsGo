package redis

import (
	"errors"
	"fmt"
	redis8 "github.com/gomodule/redigo/redis"
	"thor/util/helper"
)

var (
	redisCluster = map[string]redis8.Conn{}
)

func NewEngine() (map[string]redis8.Conn, error) {
	var err error
	redisCluster = make(map[string]redis8.Conn)

	redisConf := helper.GetAppConf().RedisConf

	for _, v := range redisConf {
		redisCluster[v.Name], err = newConnect(v)
		if err != nil {
			return nil, err
		}
	}

	return redisCluster, nil
}

func newConnect(r helper.RedisConf) (redis8.Conn, error) {
	addr := fmt.Sprintf("%s:%d", r.Host, r.Port)

	db, err := redis8.Dial(r.Protocol, addr)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("init redis error!")
	}

	return db, nil
}
