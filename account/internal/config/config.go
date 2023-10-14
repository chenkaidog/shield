package config

import (
	"os"
	"shield/common/logs"

	"gopkg.in/yaml.v3"
)

var serverConf ServerConfig

func Init() {
	content, err := os.ReadFile("./conf/deploy.yml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(content, &serverConf); err != nil {
		panic(err)
	}

	logs.Debug("conf: %+v", serverConf)
}

func GetMySQLConf() MySQLConf {
	return serverConf.MySQL
}

func GetRedisConf() RedisConf {
	return serverConf.Redis
}

type ServerConfig struct {
	MySQL MySQLConf `yaml:"mysql"`
	Redis RedisConf `yaml:"redis"`
}

type MySQLConf struct {
	DBName   string `yaml:"db_name"`
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConf struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
