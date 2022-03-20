package content

import (
	"testing"
	"thor/util/helper"
	"thor/util/mysql"
)

func getContentObj() *ContentModel {

	helper.LoadConf("/Users/yaha/Data/code/flyaha/EatAcomic/Thor/conf/app.toml")

	// 此处可以注册三方服务
	if err := mysql.Pr.Register(); err != nil {
		panic(err)
	}

	contentObj := NewContentModel()

	return contentObj
}

func TestContentModel_GetListById(t *testing.T) {
	cObj := getContentObj()

	l, err := cObj.GetListById(1)
	if err != nil {
		t.Fatalf("err:%s", err)
		return
	}

	for _, v := range l {
		t.Logf("uid=%d, id=%d, text=%s", v.Uid, v.Id, v.ContentText)
	}
}
