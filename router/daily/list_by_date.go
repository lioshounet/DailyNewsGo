package daily

import (
	"github.com/gin-gonic/gin"
	"thor/service/page"
	"thor/util/helper"
	"thor/util/response"
)

type ListByDateParam struct {
	Date uint64 `json:"date"`
}

func ListByDate(c *gin.Context) {

	var params ListByDateParam
	l := helper.NewLog()

	err := c.BindJSON(&params)
	if err != nil {
		l.Warnf("[list_by_date params error][err:%v]", err)
		response.OutPutError(c, response.ParamError, "参数不完整")
		return
	}
	dailyLogic := page.NewDailyLogic(l)
	output := make(map[string]interface{})
	result, err := dailyLogic.GetListByDate(params.Date)
	l.Infof("[list_by_date response output][res:%+v][err:%v]", result, err)
	if err != nil {
		response.OutPutError(c, response.ParamError, err.Error())
		return
	}

	output["list"] = result

	response.OutPutSuccess(c, output)
}
