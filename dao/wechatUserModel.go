package dao

import (
	"rpc/config"
	"time"
)

type WechatUserInterfaceR interface {
	GetById(id int64) *WechatUserModel
	GetByWidAndOpenid(wid int64, openid string) *WechatUserModel
	LimitByWid(int64, int, int) []*WechatUserModel
	Count(*WechatUserModel) int64
}

type WechatUserInterfaceW interface {
	Insert(*WechatUserModel) (int64, error)
	Update(*WechatUserModel) (int64, error)
	DeleteById(int64) bool
}

type WechatUserModel struct {
	Id         int64     `xorm:"pk Int"`
	Wid        int64     `xorm:"wid"`
	UserId     int64     `xorm:"user_id"`
	Openid     string    `xorm:"varchar(64)"`
	Nickname   string    `xorm:"varchar(64)"`
	Sex        int8      `xorm:"sex"`
	Province   string    `xorm:"varchar(20)"`
	City       string    `xorm:"varchar(20)"`
	Country    string    `xorm:"varchar(20)"`
	Language   string    `xorm:"varchar(20)"`
	Headimgurl string    `xorm:"varchar(200)"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

func (this *WechatUserModel) TableName() string {
	return "wechat_users"
}

func (this *WechatUserModel) GetById(id int64) *WechatUserModel {
	if id < 1 {
		return nil
	}
	user := new(WechatUserModel)
	user.Id = id
	has, _ := config.GetDbR(APP_DB_READ).Get(user)
	if has == false {
		return nil
	}
	return user
}

func (wu *WechatUserModel) GetByWidAndOpenid(wid int64, openid string) *WechatUserModel {
	if openid == "" || wid < 1 {
		return nil
	}
	user := new(WechatUserModel)
	user.Wid = wid
	user.Openid = openid
	has, err := config.GetDbR(APP_DB_READ).Get(user)
	if err != nil {
		config.Logger().Errorf("wechatUserModel: GetByWidAndOpenid user: %v, err: %v", user, err)
		return nil
	} else if has == false {
		return nil
	}
	return user
}

func (wu *WechatUserModel) LimitByWid(wid int64, index int, limit int) (users []*WechatUserModel) {
	err := config.GetDbR(APP_DB_READ).Where("wid = ?", wid).Limit(limit, (index-1)*limit).Find(&users)
	if err != nil {
		return nil
	}
	return users
}

func (this *WechatUserModel) Count(user *WechatUserModel) int64 {
	total, err := config.GetDbW(APP_DB_WRITE).Count(user)
	if err != nil {
		return 0
	}
	return total
}

func (wu *WechatUserModel) Insert(user *WechatUserModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(user)
}

func (wu *WechatUserModel) Update(user *WechatUserModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).Id(user.Id).Update(user)
}

func (wu *WechatUserModel) DeleteById(id int64) bool {
	if id == 0 {
		return false
	}
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&WechatUserModel{})
	if err != nil {
		return false
	}
	return true
}
