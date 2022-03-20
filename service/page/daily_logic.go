package page

import (
	"errors"
	"github.com/sirupsen/logrus"
	"thor/service/data/mysql/content"
	"thor/service/data/mysql/user"
)

type DailyLogic struct {
	dsObj     *content.ContentModel
	userDsObj *user.UserModel
	logger    *logrus.Logger
}

type ContentUpdateRes struct {
	Row uint64 `json:"row"`
}

type ContentListRes struct {
	Id            uint64 `json:"id"`
	ContentTagStr string `json:"content_tag_str"`
	IdentityText  string `json:"identity_text"`
	ContentTag    uint64 `json:"content_tag"`
	ContentText   string `json:"content_text"`
	Uid           uint64 `json:"uid"`
	UserName      string `json:"user_name"`
	MTime         uint64 `json:"m_time"`
	Avatar        string `json:"avatar"`
}

func NewDailyLogic(l *logrus.Logger) *DailyLogic {
	return &DailyLogic{
		dsObj:     content.NewContentModel(),
		userDsObj: user.NewUserModel(),
		logger:    l,
	}
}

func (t *DailyLogic) Create(uid, contentTag uint64, contentText string) (*ContentUpdateRes, error) {
	if uid < 1 || contentTag < 1 || contentText == "" {
		t.logger.Warnf("[params error][uid:%d][content_tag:%d][content_text:%s]", uid, contentTag, contentText)
		return nil, errors.New("params error")
	}

	row, err := t.dsObj.InsertContent(uid, contentTag, content.TypeDaily, contentText)
	if err != nil {
		t.logger.Warnf("[insert content error][uid:%d][content_tag:%d][content_text:%s][err:%v]", uid, contentTag, contentText, err)
		return nil, err
	}

	return &ContentUpdateRes{
		Row: row,
	}, nil
}

func (t *DailyLogic) GetRecordById(id uint64) (*ContentListRes, error) {
	if id < 1 {
		t.logger.Warnf("[params error][id:%d]]", id)
		return nil, errors.New("params error")
	}

	result, err := t.dsObj.GetListById(id)
	if err != nil {
		t.logger.Warnf("[get content list by id error][id:%d][err:%v]", id, err)
		return nil, err
	}

	for _, v := range result {

		return &ContentListRes{
			Id:          v.Id,
			ContentTag:  v.ContentTag,
			ContentText: v.ContentText,
			Uid:         v.Uid,
			MTime:       v.MTime,
		}, nil
	}

	return nil, nil
}

func (t *DailyLogic) GetList() ([]*ContentListRes, error) {
	result, err := t.dsObj.GetAllContentList(content.TypeDaily, 0, 1000)
	if err != nil {
		t.logger.Warnf("[get content list error][err:%v]", err)
		return nil, err
	}

	if len(result) < 1 {
		t.logger.Warnf("[get content list empty][err:%v]", err)
		return nil, nil
	}

	// 获取所有uid
	uids := []uint64{}
	for _, v := range result {
		uids = append(uids, v.Uid)
	}

	userList, err := t.userDsObj.GetUserListByUids(uids)
	if err != nil {
		t.logger.Warnf("[get user list by uids error][uids:%v][err:%v]", uids, err)
		return nil, err
	}

	userMap := make(map[uint64]*user.User)
	for _, v := range userList {
		userMap[v.Uid] = v
	}

	list := make([]*ContentListRes, 0)

	for _, v := range result {
		u, ok := userMap[v.Uid]
		if !ok {
			continue
		}

		identityStr, ok := user.IdentityMap[u.Identity]
		if !ok {
			identityStr = user.IdentityMap[user.OTHER]
		}

		if u.Avatar == "" {
			u.Avatar = "https://img0.baidu.com/it/u=3311900507,1448170316&fm=26&fmt=auto&gp=0.jpg"
		}

		l := ContentListRes{
			Id:           v.Id,
			IdentityText: identityStr,
			ContentTag:   v.ContentTag,
			ContentText:  v.ContentText,
			Uid:          v.Uid,
			MTime:        v.MTime,
			UserName:     u.UserName,
			Avatar:       u.Avatar,
		}

		list = append(list, &l)
	}

	return list, nil

}

func (t *DailyLogic) GetListByDate(date uint64) ([]*ContentListRes, error) {
	if date < 0 {
		t.logger.Warnf("[params error][date:%d]", date)
		return nil, errors.New("params error")
	}

	result, err := t.dsObj.GetContentListByDate(date, content.TypeDaily, 0, 1000)
	if err != nil {
		t.logger.Warnf("[get content list by date error][date:%d][err:%v]", date, err)
		return nil, err
	}

	if len(result) < 1 {
		t.logger.Warnf("[get content list by date empty][date:%d][err:%v]", date, err)
		return nil, nil
	}

	// 获取所有uid
	uids := []uint64{}
	for _, v := range result {
		uids = append(uids, v.Uid)
	}

	userList, err := t.userDsObj.GetUserListByUids(uids)
	if err != nil {
		t.logger.Warnf("[get user list by uids error][uids:%v][err:%v]", uids, err)
		return nil, err
	}

	userMap := make(map[uint64]*user.User)
	for _, v := range userList {
		userMap[v.Uid] = v
	}

	list := make([]*ContentListRes, 0)

	for _, v := range result {
		u, ok := userMap[v.Uid]
		if !ok {
			continue
		}

		if u.Avatar == "" {
			u.Avatar = "https://img0.baidu.com/it/u=3311900507,1448170316&fm=26&fmt=auto&gp=0.jpg"
		}

		identityStr, ok := user.IdentityMap[u.Identity]
		if !ok {
			identityStr = user.IdentityMap[user.OTHER]
		}

		l := ContentListRes{
			Id:          v.Id,
			ContentTag:  v.ContentTag,
			IdentityText: identityStr,
			ContentText: v.ContentText,
			Uid:         v.Uid,
			MTime:       v.MTime,
			UserName:    u.UserName,
			Avatar:      u.Avatar,
		}

		list = append(list, &l)
	}

	return list, nil
}
