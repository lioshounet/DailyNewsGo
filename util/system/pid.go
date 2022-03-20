package system

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// 阻塞等待程序内部的Stop通道信号
func WaitStop() {
	<-srv.stop
}

// 获取pid位置
func (opts *Options) GetPidFile(appName string) string {
	return strings.TrimRight(opts.PidDir, "/") + "/" + appName + ".pid"
}

// 读取文件的进程号
func ReadPidFile(path string) (int, error) {
	fd, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer fd.Close()

	buf := bufio.NewReader(fd)
	line, err := buf.ReadString('\n')
	if err != nil {
		return -1, err
	}
	line = strings.TrimSpace(line)
	return strconv.Atoi(line)
}

// 将进程号写入文件
func WritePidFile(path string, pidArgs ...int) error {
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	var pid int
	if len(pidArgs) > 0 {
		pid = pidArgs[0]
	} else {
		pid = os.Getpid()
	}
	_, err = fd.WriteString(fmt.Sprintf("%d\n", pid))
	return err
}

// 关闭服务
func CloseService() {
	if srv.debug {
		fmt.Println("close service")
	}
	Free()
}

// 监听信号量
func RegisterSignal() {
	go func() {
		var sigs = []os.Signal{
			syscall.SIGHUP,
			syscall.SIGUSR1,
			syscall.SIGUSR2,
			syscall.SIGINT,
			syscall.SIGTERM,
		}
		c := make(chan os.Signal)
		signal.Notify(c, sigs...)
		for {
			sig := <-c // blocked
			HandleSignal(sig)
		}
	}()
}

// 处理进程的信号量
func HandleSignal(sig os.Signal) {
	switch sig {
	case syscall.SIGINT:
		fallthrough
	case syscall.SIGTERM:
		Stop()
	default:
	}
}