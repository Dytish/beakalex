package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

var config *AppConfig
var once sync.Once

type AppConfig struct {
	isDebug bool `yaml:"is_debug"`

	StorageDB struct {
		Host     string `yaml:"host"`     // ip address
		Port     string `yaml:"port"`     // port
		Username string `yaml:"username"` // database username
		Database string `yaml:"database"` // database name
		Password string `yaml:"password"` // password
	} `yaml:"storageDB"`
}

func GetConfig() *AppConfig {
	return config
}

func init() {
	once.Do(func() {
		config = &AppConfig{}
		err := cleanenv.ReadConfig("./config.yaml", config)
		if err != nil {
			panic(err)
		}
	})

}
