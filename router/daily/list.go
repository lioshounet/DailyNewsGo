package daily

import (
	"github.com/gin-gonic/gin"
	"thor/service/page"
	"thor/util/helper"
	"thor/util/response"
)

func List(c *gin.Context) {

	l := helper.NewLog()

	dailyLogic := page.NewDailyLogic(l)
	output := make(map[string]interface{})
	result, err := dailyLogic.GetList()
	l.Infof("[list_by_date response output][res:%+v][err:%v]", result, err)
	if err != nil {
		response.OutPutError(c, response.ParamError, err.Error())
		return
	}

	output["list"] = result

	response.OutPutSuccess(c, output)
}
