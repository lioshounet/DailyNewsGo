package helper

import (
	"github.com/BurntSushi/toml"
	"os"
	"thor/util/system"
)

const (
	ProdEnv = "production"
	AppName = ""
	Version = "v1.0.0"
)

type RedisConf struct {
	Name     string `toml:"name"`
	Host     string `toml:"host"`
	Port     uint64 `toml:"port"`
	Passwd   string `toml:"passwd"`
	DB       uint64 `toml:"db"`
	Protocol string `toml:"protocol"`
}

type MysqlConfig struct {
	DBName   string `toml:"dbname"`
	Port     int    `toml:"port"`
	Password string `toml:"passwd"`
	UserName string `toml:"user"`
	HostIP   string `toml:"host"`
}

type LdapConf struct {
	Name        string `toml:"name"`
	Host        string `toml:"host"`
	Port        uint64 `toml:"port"`
	Base        string `toml:"base"`
	Dn          string `toml:"bind_dn"`
	UseSSL      bool   `toml:"use_ssl"`
	SkipTLS     bool   `toml:"skip_tls"`
	Password    string `toml:"password"`
	UserFilter  string `toml:"user_filter"`
	GroupFilter string `toml:"group_filter"`
}

type GinConf struct {
	Port string `toml:"port"`
}

type LogConf struct {
	Path string `toml:"path"`
	Name string `toml:"name"`
}

type App struct {
	AppName   string        `toml:"app_name"`
	Env       string        `toml:"env"`
	Debug     bool          `toml:"debug"`
	Log       LogConf       `toml:"log"`
	GinConf   GinConf       `toml:"gin"`
	MysqlConf []MysqlConfig `toml:"mysql"`
	RedisConf []RedisConf   `toml:"redis"`
	LdapConf  LdapConf      `toml:"ldap"`
}

var app *App

func LoadConf(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if _, err := toml.DecodeFile(filePath, &app); err != nil {
		return err
	}

	system.SetDebug(app.Debug)

	return nil
}

func GetAppConf() *App {
	return app
}
