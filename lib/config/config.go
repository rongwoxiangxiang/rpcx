package config

import (
	"github.com/pelletier/go-toml"
	"log"
)

var (
	Conf *toml.Tree
)

func InitConfig() {
	var err error
	Conf, err = toml.LoadFile("./application.toml")
	if err != nil {
		log.Fatal("Init application config error ", err.Error())
	}
}
