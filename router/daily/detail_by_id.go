package daily

import (
	"github.com/gin-gonic/gin"
	"thor/service/page"
	"thor/util/helper"
	"thor/util/response"
)

type DetailByIdParam struct {
	Id uint64 `json:"id"`
}

func DetailById(c *gin.Context) {

	var params DetailByIdParam
	l := helper.NewLog()

	err := c.BindJSON(&params)
	if err != nil {
		l.Warnf("[detail_by_id params error][err:%v]", err)
		response.OutPutError(c, response.ParamError, "参数不完整")
		return
	}

	dailyLogic := page.NewDailyLogic(l)
	output, err := dailyLogic.GetRecordById(params.Id)
	l.Infof("[detail_by_id response output][res:%+v][err:%v]", output, err)
	if err != nil {
		response.OutPutError(c, response.ParamError, err.Error())
		return
	}

	response.OutPutSuccess(c, output)
}
