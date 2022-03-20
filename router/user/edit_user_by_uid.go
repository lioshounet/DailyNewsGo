package user

import (
	"github.com/gin-gonic/gin"
	"thor/service/page"
	"thor/util/helper"
	"thor/util/response"
)

type EditUserByUidParam struct {
	Uid      uint64 `json:"uid"`
	Phone    string `json:"phone"`
	Identity uint64 `json:"identity"`
}

func EditUserByUid(c *gin.Context) {
	var params EditUserByUidParam

	l := helper.NewLog()

	err := c.BindJSON(&params)
	if err != nil {
		l.Warnf("[login params error][err:%v]", err)
		response.OutPutError(c, response.ParamError, "用户名或密码不存在")
		return
	}

	userLogic := page.NewUserLogic(l)
	output, err := userLogic.EditIdentityWithPhoneByUid(params.Uid, params.Identity, params.Phone)
	l.Infof("[login response output][res:%+v][err:%v]", output, err)
	if err != nil {
		response.OutPutError(c, response.ParamError, err.Error())
		return
	}

	response.OutPutSuccess(c, output)
}
