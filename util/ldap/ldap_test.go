package ldap

import (
	"github.com/jtblin/go-ldap-client"
	"testing"
)

func TestLdap(t *testing.T) {
	ldapConf := LdapConf{
		Name:        "blade",
		Base:        "dc=flyaha,dc=top",
		Host:        "127.0.0.1",
		Port:        389,
		UseSSL:      false,
		SkipTLS:     true,
		Dn:          "cn=admin,dc=flyaha,dc=top",
		Password:    "test",
		UserFilter:  "(uid=%s)",
		GroupFilter: "(memberUid=%s)",
	}

	client, err := NewEngine(ldapConf)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	var ldapSev *ldap.LDAPClient

	if l, ok := client["blade"]; ok {
		ldapSev = l
	}

	username := "yangzuhao"
	password := "123456"

	ok, user, err := ldapSev.Authenticate(username, password)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	if !ok {
		t.Error(err)
		t.Failed()
	}

	t.Log(user)
}
