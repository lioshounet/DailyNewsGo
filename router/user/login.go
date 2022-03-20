package user

import (
	"github.com/gin-gonic/gin"
	"thor/service/page"
	"thor/util/helper"
	"thor/util/response"
)

type LoginParam struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func Login(c *gin.Context) {
	var params LoginParam

	l := helper.NewLog()

	err := c.BindJSON(&params)
	if err != nil {
		l.Warnf("[login params error][err:%v]", err)
		response.OutPutError(c, response.ParamError, "用户名或密码不存在")
		return
	}

	userLogic := page.NewUserLogic(l)
	output, err := userLogic.Login(params.UserName, params.PassWord)
	l.Infof("[login response output][res:%+v][err:%v]", output, err)
	if err != nil {
		response.OutPutError(c, response.ParamError, err.Error())
		return
	}

	response.OutPutSuccess(c, output)
}
