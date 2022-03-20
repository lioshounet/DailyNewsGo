package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 成功时返回
func OutPutSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":        Success,
		"message":     "ok",
		"request_uri": c.Request.URL.Path,
		"data":        data,
	})
	c.Abort()
}

// 失败时返回
func OutPutError(c *gin.Context, code int, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = GetMsg(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        code,
		"message":     message,
		"request_uri": c.Request.URL.Path,
		"data":        make(map[string]string),
	})
	c.Abort()
}
