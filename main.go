package main

import (
	"errors"
	"fmt"
	"os"
	"thor/bootstarp"
	"thor/util/helper"
	"thor/util/system"
)

func main() {
	opts := option()

	handleCmd(opts)

	run(opts)
}

// 解析启动命令
func option() *system.Options {
	opts := system.GetOptions()

	if opts.Version {
		fmt.Printf("[%v] %v\n", helper.AppName, helper.Version)
		os.Exit(0)
	}

	return opts
}

// 发送命令
func handleCmd(opts *system.Options) {
	if opts.Cmd != "" {
		pidFile := opts.GetPidFile(helper.AppName)
		err := system.HandleUserCmd(opts.Cmd, pidFile)
		if err != nil {
			fmt.Printf("执行命令（%s）失败 %s\n", opts.Cmd, err)
		} else {
			fmt.Printf("执行命令（%s）成功 \n ", opts.Cmd)
		}
		os.Exit(0)
	}
}

func run(opts *system.Options) {
	// 根据启动命令行参数，决定启动哪种服务模式
	switch opts.AppType {
	case "web":
		bootstarp.Run(opts)
	default:
		fmt.Printf("-a 参数不存在 %v", opts.AppType)
		errors.New("程序启动失败")
	}
	return
}
