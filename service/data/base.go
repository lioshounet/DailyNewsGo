package data

import (
	"database/sql"
	redis8 "github.com/gomodule/redigo/redis"
	ldap2 "github.com/jtblin/go-ldap-client"
	"thor/util/ldap"
	"thor/util/mysql"
	"thor/util/redis"
)

const BladeMysqlName = "work_thor"
const BladeRedisName = "thor_cache"
const BladeLdapName = "thor_ldap"

func GetBladeDb() *sql.DB {
	db := mysql.GetDb(BladeMysqlName)
	return db
}

func GetBladeRedis() redis8.Conn {
	cache := redis.GetDb(BladeRedisName)
	return cache
}

func GetBladeLdap() *ldap2.LDAPClient {
	l := ldap.GetDb(BladeLdapName)
	return l
}
