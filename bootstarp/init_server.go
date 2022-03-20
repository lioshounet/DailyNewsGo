package bootstarp

import (
	"thor/util/ldap"
	"thor/util/mysql"
	"thor/util/redis"
	"thor/util/system"
)

// 全局变量
// var App *container.Container

func initServer() error {
	// App = container.App

	// 此处可以注册三方服务
	if err := mysql.Pr.Register(); err != nil {
		return err
	}

	if err := redis.Pr.Register(); err != nil {
		return err
	}

	if err := ldap.Pr.Register(); err != nil {
		return err
	}

	// 注册应用停止时调用
	system.MultiRegister()

	return nil
}
