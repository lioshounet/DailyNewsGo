package user

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"thor/service/data"
	"thor/util/helper"
	"time"
)

type UserModel struct {
	Db *sql.DB
	// Logger *log.Logger
}

// func NewUserModel(logger *log.Logger) *UserModel {
func NewUserModel() *UserModel {
	return &UserModel{
		// Logger: logger,
		Db: data.GetBladeDb(),
	}
}

func (t *UserModel) getTableName() string {
	return "user"
}

func (t *UserModel) CreateUser(um *User) (uint64, error) {
	if um == nil {
		return 0, errors.New("params fail!")
	}

	nowTime := time.Now().Unix()

	uid, err := helper.GenUid()
	if err != nil {
		return 0, err
		// t.Logger.Warnf("gen uid fail err:%s", err)
	}

	if um.UserName == "" || um.Email == "" || um.Sn == "" || um.GivenName == "" {
		return 0, errors.New("params fail")
	}

	insertFmt := "insert into %s (%s) values (?,?,?,?,?,?,?,?,?,?) on duplicate key update user_name=?,email=?,sn=?,given_name=?,m_time=?"
	execSql := fmt.Sprintf(insertFmt, t.getTableName(), insertField)

	// t.Logger.Infof("GetUserByUid sql:%s", execSql)

	result, err := t.Db.Exec(execSql, uid, um.UserName, um.Identity, um.Email, um.Phone, um.Avatar, um.Sn, um.GivenName, nowTime, nowTime, um.UserName, um.Email, um.Sn, um.GivenName, nowTime)

	if err != nil {
		return 0, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(row), nil

}

func (t *UserModel) GetUserListByUids(uids []uint64) ([]*User, error) {
	if len(uids) < 1 {
		return nil, errors.New("params errors")
	}

	uids_str := []string{}
	for _, v := range uids {
		uids_str = append(uids_str, strconv.FormatInt(int64(v), 10))
	}

	uid_str := strings.Join(uids_str, ",")

	sqlFmt := "select %s from user where uid in (%s) and deleted = %d"
	querySql := fmt.Sprintf(sqlFmt, selectField, uid_str, DeletedOff)
	user, err := t.userQuery(querySql)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (t *UserModel) EditIdentityWithPhoneByUid(uid, identity uint64, phone string) (uint64, error) {
	if uid < 1 || identity < 1 {
		return 0, errors.New("params errors")
	}

	if phone == "" {
		return 0, errors.New("params errors")
	}

	nowTime := time.Now().Unix()

	updateFmt := "update %s set identity=%d, phone=?, m_time=%d where uid=%d"
	execSql := fmt.Sprintf(updateFmt, t.getTableName(), identity, nowTime, uid)

	result, err := t.Db.Exec(execSql, phone)

	if err != nil {
		return 0, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(row), nil
}

func (t *UserModel) GetUserRecordByUserName(username string) (*User, error) {

	if username == "" {
		return nil, errors.New("prams errors, uid < 1")
	}

	sqlFmt := "select %s from user where user_name = '%s' and deleted = %d"
	querySql := fmt.Sprintf(sqlFmt, selectField, username, DeletedOff)

	user, err := t.userQuery(querySql)
	if err != nil {
		return nil, err
	}

	if len(user) < 1 {
		return nil, errors.New("get user info empty")
	}

	return user[0], nil
}

func (t *UserModel) GetUserRecordByUid(uid uint64) (*User, error) {

	if uid < 1 {
		return nil, errors.New("prams errors, uid < 1")
	}

	sqlFmt := "select %s from user where uid = %d and deleted = %d"
	querySql := fmt.Sprintf(sqlFmt, selectField, uid, DeletedOff)

	user, err := t.userQuery(querySql)
	if err != nil {
		return nil, err
	}

	if len(user) < 1 {
		return nil, errors.New("get user info empty")
	}

	return user[0], nil
}

func (t *UserModel) userQuery(querySql string) ([]*User, error) {

	userList := make([]*User, 0)
	rows, err := t.Db.Query(querySql)
	if err != nil {
		return nil, err
	}

	var id, uid, identity, deleted, c_time, m_time uint64
	var user_name, phone, email, avatar, sn, given_name string

	for rows.Next() {
		if err := rows.Scan(&id, &uid, &user_name, &identity, &email, &phone, &avatar, &sn, &given_name, &deleted, &c_time, &m_time); err != nil {
			return nil, err
		}

		user := &User{
			Id:        id,
			Uid:       uid,
			UserName:  user_name,
			Identity:  identity,
			Email:     email,
			Phone:     phone,
			Avatar:    avatar,
			Sn:        sn,
			GivenName: given_name,
			Deleted:   deleted,
			CTime:     c_time,
			MTime:     m_time,
		}

		userList = append(userList, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userList, nil
}
