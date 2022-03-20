package ldap

import (
	"github.com/jtblin/go-ldap-client"
	"sync"
	"thor/util/container"
)

const (
	SingletonMain = "ldap"
)

var Pr *provider

type provider struct {
	mu   sync.RWMutex
	mp   map[string]interface{}
	name string
}

func init() {
	Pr = new(provider)
	Pr.mp = make(map[string]interface{})
}

// 将服务注册到容器中
func (p *provider) Register() error {

	p.mu.Lock()
	p.mp[SingletonMain] = SingletonMain
	if len(p.mp) == 1 {
		p.name = SingletonMain
	}
	p.mu.Unlock()

	if _, err := setSingleton(); err != nil {
		return err
	}

	return nil
}

// 获取单利
func getSingleton(name string) *ldap.LDAPClient {
	rc := container.App.GetSingleton(SingletonMain)

	if rc == nil {
		return nil
	}

	cluster := rc.(map[string]*ldap.LDAPClient)

	if l, ok := cluster[name]; ok {
		return l
	}

	return nil
}

func setSingleton() (map[string]*ldap.LDAPClient, error) {
	l, err := NewEngine()
	if err == nil {
		container.App.SetSingleton(SingletonMain, l)
	}
	return l, nil
}

// 获取容器对象
func GetDb(name string) *ldap.LDAPClient {
	return getSingleton(name)
}
