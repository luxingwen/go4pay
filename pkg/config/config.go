package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port                       int
	Dsn                        string    `json:"dsn"`
	Log                        LogConfig `json:"log"`
	SecretKeyOnOpenpixPlatform string    `json:"secretKeyOnOpenpixPlatform"`
	IsSign                     bool      `json:"isSign"`
}

type LogConfig struct {
	Dir      string `json:"dir"`
	Filename string `json:"filename"`
}

var config Config

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &config); err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return config
}

func GetLogConfig() LogConfig {
	return config.Log
}
