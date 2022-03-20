package ldap

import (
	"github.com/jtblin/go-ldap-client"
	"thor/util/helper"
)

var (
	ldapServer map[string]*ldap.LDAPClient
)

func NewEngine() (map[string]*ldap.LDAPClient, error) {
	var err error

	ldapServer = make(map[string]*ldap.LDAPClient)

	ldapConfig := helper.GetAppConf().LdapConf

	client, err := newConnect(ldapConfig)

	if err != nil {
		return nil, err
	}

	ldapServer[ldapConfig.Name] = client

	return ldapServer, nil
}

func newConnect(config helper.LdapConf) (*ldap.LDAPClient, error) {

	attributes := []string{"givenName", "sn", "mail", "uid"}

	client := &ldap.LDAPClient{
		Base:         config.Base,
		Host:         config.Host,
		Port:         int(config.Port),
		UseSSL:       config.UseSSL,
		SkipTLS:      config.SkipTLS,
		BindDN:       config.Dn,
		BindPassword: config.Password,
		UserFilter:   config.UserFilter,
		GroupFilter:  config.GroupFilter,
		Attributes:   attributes,
	}

	return client, nil
}
