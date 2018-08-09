package util

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	Mongo Mongo  `yaml:"mongo"`
}
type Mongo struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB   string `yaml:"db"`
}

var config *Config

func InitConfig() *Config {
	configFile, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatal("read config yaml error: %v", err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("unmarshal config error: %v", err)
	}
	return config
}

func GetConfig() *Config {
	return config
}
