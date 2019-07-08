package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pelletier/go-toml"
	"log"
	"strings"
)

var (
	debug    bool
	enginesR map[string]*xorm.Engine
	enginesW map[string]*xorm.Engine
)

func SetDbDebug(debug bool) {
	debug = debug
}

func InitStoreDb() {
	enginesR = make(map[string]*xorm.Engine)
	enginesW = make(map[string]*xorm.Engine)
	sources := Conf.Get("source")
	if sources == nil {
		log.Fatal("Init orm failed to initialized: database source null")
	}
	//初始化
	var engine *xorm.Engine
	for enginesName, source := range sources.(*toml.Tree).ToMap() {
		engine = initMysqlDb(source.(string))
		if engine == nil {
			continue
		}
		if strings.Contains(enginesName, "write") {
			enginesW[enginesName] = engine
		} else {
			enginesR[enginesName] = engine
		}
		log.Printf("Init database %s SUCCESS!", enginesName)
	}
	log.Println("InitMysqlDb SUCCESS!")
}

func initMysqlDb(source string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", source)
	engine.ShowSQL(debug)
	if err != nil {
		log.Println("orm failed to initialized: %v", err)
		return nil
	}
	return engine
}

func GetDbR(store string) *xorm.Engine {
	if engine, ok := enginesR[store]; ok == true {
		return engine
	}
	return nil
}

func GetDbW(store string) *xorm.Engine {
	if engine, ok := enginesW[store]; ok == true {
		return engine
	}
	return nil
}

func GetDefaultR() *xorm.Engine {
	if engine, ok := enginesR["default_read"]; ok == true {
		return engine
	}
	return nil
}
func GetDefaultW() *xorm.Engine {
	if engine, ok := enginesW["default_write"]; ok == true {
		return engine
	}
	return nil
}
