package dao

import (
	"rpc/common"
	"rpc/config"
	"time"
)

type WechatUserInterfaceR interface {
	GetById(id int64) (*WechatUserModel, error)
	GetByWidAndOpenid(wid int64, openid string) (*WechatUserModel, error)
	LimitUnderWidList(index int, limit int) ([]*WechatUserModel, error)
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
	Sex        int       `xorm:"sex"`
	Province   string    `xorm:"varchar(20)"`
	City       string    `xorm:"varchar(20)"`
	Country    string    `xorm:"varchar(20)"`
	Language   string    `xorm:"varchar(20)"`
	Headimgurl string    `xorm:"varchar(200)"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

func (wu *WechatUserModel) TableName() string {
	return "wechat_users"
}

func (this *WechatUserModel) GetById(id int64) (*WechatUserModel, error) {
	if id < 1 {
		return nil, common.ErrDataGet
	}
	user := new(WechatUserModel)
	user.Id = id
	has, err := config.GetDbR(APP_DB_READ).Get(user)
	if err != nil {
		err = common.ErrDataGet
	} else if has == false {
		err = common.ErrDataEmpty
	}
	return user, err
}

func (wu *WechatUserModel) GetByWidAndOpenid(wid int64, openid string) (*WechatUserModel, error) {
	if openid == "" || wid < 1 {
		return nil, common.ErrDataGet
	}
	user := new(WechatUserModel)
	user.Wid = wid
	user.Openid = openid
	has, err := config.GetDbR(APP_DB_READ).Get(user)
	if err != nil {
		return nil, common.ErrDataGet
	} else if has == false {
		return nil, common.ErrDataEmpty
	}
	return user, err
}

func (wu *WechatUserModel) LimitUnderWidList(index int, limit int) (users []*WechatUserModel, err error) {
	err = config.GetDbR(APP_DB_READ).Where("wid = ?", wu.Wid).Limit(limit, (index-1)*limit).Find(&users)
	if err != nil {
		return nil, common.ErrDataFind
	}
	return users, nil

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
