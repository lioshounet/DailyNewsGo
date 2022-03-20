package redis

import (
	redis8 "github.com/gomodule/redigo/redis"
	"os"
	"testing"
	"thor/util/helper"
)

func TestRedis(t *testing.T) {

	err := helper.LoadConf("/Users/yaha/Data/code/flyaha/EatAcomic/Thor/conf/app.toml")
	if err != nil {
		t.Error(err)
		os.Exit(-1)
	}
	rdb, err := NewEngine()
	if err != nil {
		t.Error(err)
		os.Exit(-1)
	}

	var client redis8.Conn

	if l, ok := rdb["thor_cache"]; ok {
		client = l
	}

	result, err := client.Do("set", "1234", "12345678")
	t.Log(result)

	result, err = client.Do("get", "1234")

	t.Log(result)
}
