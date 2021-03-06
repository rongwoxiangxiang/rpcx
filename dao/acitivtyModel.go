package dao

import (
	"rpc/config"
	"time"
)

type ActivityInterfaceR interface {
	GetById(int64) *ActivityModel
	ListByWid(int64, int, int) []*ActivityModel
	Count(*ActivityModel) int64
}

type ActivityInterfaceW interface {
	Insert(*ActivityModel) (int64, error)
	Update(*ActivityModel) (int64, error)
	DeleteById(int64) bool
}

type ActivityModel struct {
	Id           int64 `xorm:"pk"`
	Wid          int64
	Name         string
	Desc         string
	ActivityType int32
	RelativeId   int64
	Extra        string
	TimeStarted  time.Time
	TimeEnd      time.Time
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}

func (this *ActivityModel) TableName() string {
	return "activities"
}

func (this *ActivityModel) GetById(id int64) *ActivityModel {
	activity := new(ActivityModel)
	activity.Id = id
	has, err := config.GetDbR(APP_DB_READ).Get(activity)
	if has != true || err != nil {
		return nil
	}
	return activity
}

func (this *ActivityModel) ListByWid(wid int64, index, limit int) (activities []*ActivityModel) {
	err := config.GetDbR(APP_DB_READ).Where("wid = ?", wid).Limit(limit, (index-1)*limit).Find(&activities)
	if err != nil {
		return nil
	}
	return activities
}

func (this *ActivityModel) Count(activity *ActivityModel) int64 {
	total, err := config.GetDbW(APP_DB_WRITE).Count(activity)
	if err != nil {
		return 0
	}
	return total
}

func (this *ActivityModel) Insert(activity *ActivityModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(activity)
}

func (this *ActivityModel) Update(activity *ActivityModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).Id(activity.Id).Update(activity)
}

func (this *ActivityModel) DeleteById(id int64) bool {
	if id < 1 {
		return false
	}
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&ActivityModel{})
	if err != nil {
		config.Logger().Error("ActivityModel DeleteById err: ", err)
		return false
	}
	return true
}
