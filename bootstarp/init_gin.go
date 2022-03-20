package bootstarp

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
	"syscall"
	"thor/util/helper"
	"thor/util/system"
	"time"
)

func initGin(pidPath string, registerRoute func(engine *gin.Engine)) error {
	engine := gin.Default()
	engine.Use(CorsMiddleware())
	engine.Use(AccessLog())
	// 注册路由
	registerRoute(engine)

	ginConfig := helper.GetAppConf().GinConf

	// 默认8080端口
	if ginConfig.Port == "" {
		ginConfig.Port = ":8080"
	}

	if err := runGin(engine, ginConfig.Port, pidPath); err != nil {
		return err
	}

	go func() {
		system.Stop()
	}()

	system.WaitStop()
	return nil
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()


		logger := helper.NewLog()

		// 日志格式
		logger.Infof("[code=%d][cost=%v][ip=%s][method=%s][uri=%s]",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func runGin(engine *gin.Engine, addr, pidPath string) error {
	// Gin的调试模式
	// if !system.GetDebug() {
	gin.SetMode(gin.ReleaseMode)
	// }

	server := endless.NewServer(addr, engine)
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		if gin.Mode() != gin.ReleaseMode {
			fmt.Printf("Actual pid is %d \n\r", pid)
		}
		system.WritePidFile(pidPath, pid)
	}
	err := server.ListenAndServe()
	return err
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
