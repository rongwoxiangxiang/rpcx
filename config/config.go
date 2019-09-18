package config

import (
	"github.com/pelletier/go-toml"
	"log"
)

var (
	debug bool
	Conf  *toml.Tree
)

func init() {
	var err error
	Conf, err = toml.LoadFile("./application.toml")
	if err != nil {
		log.Fatal("Init application config error ", err.Error())
	}
	debug = Conf.Get("application.debug").(bool)
	initLogger()
	initMysql()
	initRedis()
}

func initMysql() {
	sources := Conf.Get("source")
	if sources != nil {
		InitStoreDb(sources.(*toml.Tree))
	}
}

func initRedis() {
	redis := Conf.Get("redis")
	if redis != nil {
		InitRedis(redis.(*toml.Tree))
	}
}

func initLogger() {
	conf := Conf.Get("logger")
	if conf != nil {
		InitLogger(conf.(*toml.Tree))
	}

}
