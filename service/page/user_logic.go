package page

import (
	"errors"
	"github.com/sirupsen/logrus"
	"thor/service/data/cache"
	"thor/service/data/ldap"
	"thor/service/data/mysql/user"
	"thor/util/helper"
)

type UserInfo struct {
	Uid      uint64 `json:"uid"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

// func GetUserInfo(userName string) (*UserInfo, error) {
//
// 	user, err := cache.GetUserInfoByUserName(userName)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &UserInfo{
// 		Uid:      user.Uid,
// 		UserName: user.UserName,
// 		Email:    user.Email,
// 	}, nil
//
// }

type UserLogic struct {
	cacheObj *cache.UserCache
	dsObj    *user.UserModel
	logger   *logrus.Logger
}

type UserUpdateRes struct {
	Row uint64 `json:"row"`
}

func NewUserLogic(l *logrus.Logger) *UserLogic {
	return &UserLogic{
		cacheObj: cache.NewUserCache(),
		dsObj:    user.NewUserModel(),
		logger:   l,
	}
}

func (t *UserLogic) EditIdentityWithPhoneByUid(uid, identity uint64, phone string) (*UserUpdateRes, error) {
	if uid < 1 || identity < 1 {
		t.logger.Warnf("[params error][uid:%d][identity:%d][phone:%s]", uid, identity, phone)
		return nil, errors.New("params error")
	}

	row, err := t.dsObj.EditIdentityWithPhoneByUid(uid, identity, phone)
	if err != nil {
		t.logger.Warnf("[edit user_info_by_uid error][uid:%d][identity:%d][phone:%s][err:%v]", uid, identity, phone, err)
		return nil, err
	}

	return &UserUpdateRes{
		Row: row,
	}, nil
}

func (t *UserLogic) setUserInfoToRedis(username string) (*UserInfo, error) {
	token, err := helper.GenerateToken(username)
	if err != nil {
		return nil, err
	}

	userDsRecord, err := t.dsObj.GetUserRecordByUserName(username)
	if err != nil {
		return nil, err
	}

	userCacheParam := &cache.User{
		Uid:      userDsRecord.Uid,
		UserName: userDsRecord.UserName,
		Email:    userDsRecord.Email,
		Phone:    userDsRecord.Phone,
		Avatar:   userDsRecord.Avatar,
		Sn:       userDsRecord.Sn,
		GiveName: userDsRecord.GivenName,
		Token:    token,
	}

	ok, err := t.cacheObj.SetUserInfoByName(userCacheParam)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("login error")
	}

	return &UserInfo{
		Uid:      userCacheParam.Uid,
		UserName: userCacheParam.UserName,
		Email:    userCacheParam.Email,
		Token:    token,
	}, nil
}

func (t *UserLogic) GetUidByName(username string) (uint64, error) {
	if username == "" {
		return 0, errors.New("params error")
	}

	uCache, err := t.cacheObj.GetUserInfoByName(username)
	if err != nil {
		return 0, err
	}

	if uCache != nil {
		return uCache.Uid, nil
	}

	u, err := t.setUserInfoToRedis(username)
	if err != nil {
		return 0, err
	}

	return u.Uid, nil
}

func (t *UserLogic) UserInfo(username string) (*UserInfo, error) {
	if username == "" {
		return nil, errors.New("params error")
	}

	uCache, err := t.cacheObj.GetUserInfoByName(username)
	if err != nil {
		return nil, err
	}

	if uCache != nil {
		return &UserInfo{
			Uid:      uCache.Uid,
			UserName: uCache.UserName,
			Email:    uCache.Email,
			Token:    uCache.Token,
		}, nil
	}

	u, err := t.setUserInfoToRedis(username)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (t *UserLogic) Login(username, password string) (*UserInfo, error) {

	if username == "" || password == "" {
		t.logger.Warnf("[params error][username=%s]", username)
		return nil, errors.New("params error")
	}

	// 通过ldap验证用户是否存在
	userLdap, err := ldap.GetUserByUidAndPassword(username, password)
	if err != nil {
		t.logger.Warnf("[ldap user not exits][username:%s][err:%v]", username, err)
		return nil, err
	}

	userDsParams := &user.User{
		UserName:  userLdap.Uid,
		Sn:        userLdap.Sn,
		Email:     userLdap.Email,
		GivenName: userLdap.GivenName,
	}
	_, err = t.dsObj.CreateUser(userDsParams)
	if err != nil {
		t.logger.Warnf("[create user error][username:%s][sn:%s][email:%s][given_name:%s][err:%v]", userLdap.Uid, userLdap.Sn, userLdap.Email, userLdap.GivenName, err)
		return nil, err
	}

	u, err := t.setUserInfoToRedis(username)
	if err != nil {
		t.logger.Warnf("[set user info cache error][username=%s][err:%v]", username, err)
		return nil, err
	}

	return u, nil
}
