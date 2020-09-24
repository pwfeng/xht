package cmd

import (
	"github.com/Unknwon/goconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var cfg *goconfig.ConfigFile
var db *gorm.DB

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	cfg, err = goconfig.LoadConfigFile(strings.Join([]string{homeDir, ".xht.ini"}, string(os.PathSeparator)))
	mysqlConfig := NewMysqlConfig(cfg)

	db, err = gorm.Open(mysql.Open(mysqlConfig.Dns()), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
}

type MysqlConfig struct {
	host     string
	port     string
	username string
	password string
}

func NewMysqlConfig(cfg *goconfig.ConfigFile) *MysqlConfig {
	config := new(MysqlConfig)
	config.host = cfg.MustValue("mysql", "host", "127.0.0.1")
	config.port = cfg.MustValue("mysql", "port", "3306")
	config.username = cfg.MustValue("mysql", "username", "root")
	config.password = cfg.MustValue("mysql", "password", "")
	return config
}

func (mc *MysqlConfig) Dns() string {
	return mc.username + ":" + mc.password + "@tcp(" + mc.host + ":" + mc.port + ")/xiaohuatong?charset=utf8mb4&parseTime=True&loc=Local"
}
