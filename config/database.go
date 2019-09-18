package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pelletier/go-toml"
	"strings"
)

var (
	enginesR map[string]*xorm.Engine
	enginesW map[string]*xorm.Engine
)

func InitStoreDb(sources *toml.Tree) {
	enginesR = make(map[string]*xorm.Engine)
	enginesW = make(map[string]*xorm.Engine)
	if sources == nil {
		logger.Fatal("Init orm failed to initialized: database source null")
	}
	//初始化
	var engine *xorm.Engine
	for enginesName, source := range sources.ToMap() {
		engine = initMysqlDb(source.(string))
		if engine == nil {
			continue
		}
		if strings.Contains(enginesName, "write") {
			enginesW[enginesName] = engine
		} else {
			enginesR[enginesName] = engine
		}
		logger.Info("Init database %s SUCCESS!", enginesName)
	}
	logger.Info("InitMysqlDb SUCCESS!")
}

func initMysqlDb(source string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", source)
	engine.ShowSQL(debug)
	if err != nil {
		logger.Errorf("orm failed to initialized: %v", err)
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
