package daily

import (
	"github.com/gin-gonic/gin"
	"thor/service/page"
	"thor/util/helper"
	"thor/util/response"
)

type CreateParam struct {
	Uid         uint64 `json:"uid"`
	ContentTag  uint64 `json:"content_tag"`
	ContentText string `json:"content_text"`
}

func Create(c *gin.Context) {

	var params CreateParam
	l := helper.NewLog()

	err := c.BindJSON(&params)
	if err != nil {
		l.Warnf("[crate params error][err:%v]", err)
		response.OutPutError(c, response.ParamError, err.Error())
		return
	}

	dailyLogic := page.NewDailyLogic(l)
	userLogic := page.NewUserLogic(l)

	// 线上环境需要检查传过来的uid和登录态的uid是否一致
	if helper.GetAppConf().Env == helper.ProdEnv {
		username, err := helper.GetUserNameByToken(c)
		if err != nil {
			l.Warnf("[get_user_name_by_token error][err:%v]", err)
			response.OutPutError(c, response.ParamError, "Token已经失效")
			return
		}

		uid, err := userLogic.GetUidByName(username)
		if err != nil {
			l.Warnf("[git_uid_by_name error][err:%v]", err)
			response.OutPutError(c, response.ParamError, "登录失效")
			return
		}

		params.Uid = uid
	}

	output, err := dailyLogic.Create(params.Uid, params.ContentTag, params.ContentText)
	l.Infof("[create response output][res:%+v][err:%v]", output, err)
	if err != nil {
		response.OutPutError(c, response.ParamError, err.Error())
		return
	}

	response.OutPutSuccess(c, output)
}
