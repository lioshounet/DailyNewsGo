package system

import (
	"fmt"
	"os"
	"syscall"
)

var srv *serverInfo

// 服务的基础信息
type serverInfo struct {
	stop  chan bool
	debug bool
}

func init() {
	srv = new(serverInfo)
	srv.stop = make(chan bool, 0)
}

// 获取启动命令配置
func GetOptions() *Options {
	if options == nil {
		options = parseOptions()
	}
	return options
}

// 执行命令
func HandleUserCmd(cmd string, pidFile string) error {
	var sig os.Signal

	switch cmd {
	case "stop":
		sig = syscall.SIGTERM
	case "restart":
		sig = syscall.SIGHUP
	default:
		return fmt.Errorf("未知命令 %s", cmd)
	}

	pid, err := ReadPidFile(pidFile)
	if err != nil {
		return err
	}

	if srv.debug {
		fmt.Printf("启动成功： %v  pid %d \n", sig, pid)
	}

	proc := new(os.Process)
	proc.Pid = pid
	return proc.Signal(sig)
}

// 停止
func Stop() {
	srv.stop <- true
}

// 设置Debug
func SetDebug(debug bool) {
	srv.debug = debug
	return
}

// 获取Debug状态
func GetDebug() bool {
	return srv.debug
}
