package dao

import (
	"rpc/common"
	"rpc/config"
	"time"
)

type WechatInterfaceR interface {
	GetById(int64) *WechatModel
	GetByAppid(string) *WechatModel
	GetByFlag(string) *WechatModel
}

type WechatInterfaceW interface {
	Insert(*WechatModel) (int64, error)
	DeleteById(int64) bool
}

type WechatModel struct {
	Id             int64 `xorm:"pk"`
	Name           string
	Appid          string
	Appsecret      string
	EncodingAesKey string
	Token          string
	Flag           string
	Type           int32
	Pass           int8
	SaveInput      int8
	NeedSaveMen    int8      //该公众号hander是否持久化
	CreatedAt      time.Time `xorm:"created"`
	UpdatedAt      time.Time `xorm:"updated"`
}

func (w *WechatModel) TableName() string {
	return "wechats"
}

func (w *WechatModel) Insert(wechat *WechatModel) (int64, error) {
	return config.GetDbW(APP_DB_READ).InsertOne(wechat)
}

func (w *WechatModel) DeleteById(id int64) bool {
	if id == 0 {
		return false
	}
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&WechatModel{})
	if err != nil {
		return false
	}
	return true
}

func (w *WechatModel) GetById(id int64) *WechatModel {
	if id != 0 {
		wechat := new(WechatModel)
		wechat.Id = id
		has, err := config.GetDbR(APP_DB_READ).Get(wechat)
		if !has || err != nil {
			config.Logger().Errorf("Wechat model: getById has[%v] or err: %v", has, err)
			return nil
		}
		return wechat
	}
	return nil
}

func (w *WechatModel) GetByAppid(appid string) *WechatModel {
	if appid == "" {
		return nil
	}
	wechat := new(WechatModel)
	has, err := config.GetDbR(APP_DB_READ).Where("appid = ?", appid).Get(wechat)
	if !has || err != nil {
		return nil
	}
	return wechat
}

func (w *WechatModel) GetByFlag(flag string) *WechatModel {
	if flag == "" {
		return nil
	}
	wechat := new(WechatModel)
	has, err := config.GetDbR(APP_DB_READ).Where("flag = ?", flag).Get(wechat)
	if !has || err != nil {
		return nil
	}
	return wechat
}

func (w *WechatModel) SetPass(bool2 bool) {
	if bool2 {
		w.Pass = common.YES_VALUE
	} else {
		w.Pass = common.NO_VALUE
	}
}
func (w *WechatModel) SetSaveInput(bool2 bool) {
	if bool2 {
		w.SaveInput = common.YES_VALUE
	} else {
		w.SaveInput = common.NO_VALUE
	}
}
