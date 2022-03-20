package system

import (
	"flag"
)

// 操作命令
var options *Options

type Options struct {
	Version  bool
	Cmd      string
	AppType  string
	ConfFile string
	PidDir   string
}

// 可执行的命令
func parseOptions() *Options {
	opts := new(Options)

	// 版本信息
	flag.BoolVar(&opts.Version, "v", false, "应用信息")
	// 应用类型
	flag.StringVar(&opts.AppType, "a", "web", "应用类型（web|cron）")
	//
	flag.StringVar(&opts.Cmd, "s", "", "status|stop|restart")
	// 文件位置
	flag.StringVar(&opts.ConfFile, "c", "conf/app.toml", "配置文件路径")
	// pid路径
	flag.StringVar(&opts.PidDir, "p", "/var/run/", "pid路径")

	flag.Parse()

	return opts
}
