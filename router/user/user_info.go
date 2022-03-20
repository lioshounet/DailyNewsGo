package user

import (
	"github.com/gin-gonic/gin"
	"thor/service/page"
	"thor/util/helper"
	"thor/util/response"
)

func UserInfo(c *gin.Context) {
	l := helper.NewLog()
	username, err := helper.GetUserNameByToken(c)
	if err != nil {
		l.Warnf("[user_info params error][err:%v]", err)
		response.OutPutError(c, response.ParamError, "Token已经失效")
		return
	}
	userLogic := page.NewUserLogic(l)
	output, err := userLogic.UserInfo(username)
	l.Infof("[user_info response output][res:%+v][err:%v]", output, err)
	if err != nil {
		response.OutPutError(c, response.ParamError, err.Error())
		return
	}

	// c.SetCookie("username", output.UserName, 10, "/", "localhost", false, true)
	response.OutPutSuccess(c, output)
}
