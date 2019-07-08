package models

import (
	"rpc/lib/config"
	"time"
)

type ActivityModel struct {
	Id           int64 `xorm:"pk"`
	Wid          int64
	Name         string
	Desc         string
	ActivityType int8
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

func (this *ActivityModel) GetById(id int64) (activity *ActivityModel) {
	if id != 0 {
		activity := new(ActivityModel)
		activity.Id = id
		has, err := config.GetDefaultR().Get(activity)
		if has != true || err != nil {
			return nil
		}
		return activity
	}
	return nil
}

func (this *ActivityModel) LimitUnderWidList(index, limit, wid int) (activities []*ActivityModel) {
	if index < 1 || wid < 1 || limit < 1 {
		return nil
	}
	err := config.GetDefaultR().Where("wid = ?", wid).Limit(limit, (index-1)*limit).Find(&activities)
	if err != nil {
		return nil
	}
	return activities
}

func (this *ActivityModel) Insert(activity *ActivityModel) (int64, error) {
	return config.GetDefaultW().InsertOne(activity)
}

func (this *ActivityModel) Update(activity *ActivityModel) (int64, error) {
	return config.GetDefaultW().Id(activity.Id).Update(activity)
}

func (this *ActivityModel) DeleteById(id int64) bool {
	if id < 1 {
		return false
	}
	_, err := config.GetDefaultW().Id(id).Unscoped().Delete(&ActivityModel{})
	if err != nil {
		return false
	}
	return true
}

func (this *ActivityModel) Count(model *ActivityModel) int64 {
	total, err := config.GetDefaultR().Count(model)
	if err != nil {
		return 0
	}
	return total
}
