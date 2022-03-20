package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"thor/util/container"
)

const (
	SingletonMain = "mysql"
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
func getSingleton(name string) *sql.DB {
	rc := container.App.GetSingleton(SingletonMain)
	if rc == nil {
		return nil
	}
	dbCluster := rc.(map[string]*sql.DB)

	if db, ok := dbCluster[name]; ok {
		return db
	}

	return nil
}

func setSingleton() (map[string]*sql.DB, error) {
	db, err := NewEngine()
	if err == nil {
		container.App.SetSingleton(SingletonMain, db)
	}
	return db, nil
}

// 获取容器对象
func GetDb(name string) *sql.DB {
	return getSingleton(name)
}
