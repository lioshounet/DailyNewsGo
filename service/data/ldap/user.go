package ldap

import (
	"errors"
	"thor/service/data"
)

type User struct {
	GivenName string `json:"given_name"`
	Sn        string `json:"sn"`
	Email     string `json:"email"`
	Uid       string `json:"uid"`
}

func GetUserByUidAndPassword(userName, password string) (*User, error) {
	if userName == "" || password == "" {
		return nil, errors.New("用户名或密码错误！")
	}

	ldapSev := data.GetBladeLdap()

	// Ldap查询用户是否正常
	var givenName, sn, email, uid string
	ok, userLdapRes, err := ldapSev.Authenticate(userName, password)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("user not exits")
	}

	if g, ok := userLdapRes["givenName"]; ok {
		givenName = g
	}

	if s, ok := userLdapRes["sn"]; ok {
		sn = s
	}

	if e, ok := userLdapRes["mail"]; ok {
		email = e
	}

	if u, ok := userLdapRes["uid"]; ok {
		uid = u
	}

	return &User{
		Uid:       uid,
		GivenName: givenName,
		Email:     email,
		Sn:        sn,
	}, nil

}
