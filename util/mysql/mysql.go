package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"thor/util/helper"
)

var (
	driver       = "mysql"
	mysqlCluster map[string]*sql.DB
)

func NewEngine() (map[string]*sql.DB, error) {
	var err error
	mysqlCluster = make(map[string]*sql.DB)

	for _, v := range helper.GetAppConf().MysqlConf {
		mysqlCluster[v.DBName], err = newConnect(v)
		if err != nil {
			return nil, err
		}
	}

	return mysqlCluster, nil
}

func newConnect(config helper.MysqlConfig) (*sql.DB, error) {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", config.UserName, config.Password, config.HostIP, config.Port, config.DBName)

	db, err := sql.Open(driver, dsn)

	if err != nil {
		return nil, err
	}

	// 设置表名和字段的映射规则：驼峰转下划线

	return db, nil
}
