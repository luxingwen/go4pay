package db

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

type Config struct {
	Dsn string `json:"dsn"`
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

func GetDB() *xorm.Engine {
	if engine != nil {
		return engine
	}

	var err error
	engine, err = xorm.NewEngine("mysql", config.Dsn)
	if err != nil {
		panic(err)
	}

	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	return engine
}
