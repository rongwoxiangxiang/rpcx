package dao

import (
	"rpc/config"
	"time"
)

type PrizeHistoryInterfaceR interface {
	GetByActivityWuId(activityId, wuid int64) *PrizeHistoryModel
	LimitByActivityList(activityId int64, index int, limit int) []*PrizeHistoryModel
	Count(*PrizeHistoryModel) int64
}

type PrizeHistoryInterfaceW interface {
	Insert(*PrizeHistoryModel) (int64, error)
	DeleteById(int64) bool
}

type PrizeHistoryModel struct {
	Id         int64 `xorm:"pk"`
	ActivityId int64
	Wuid       int64
	Prize      string
	Code       string
	Level      string
	CreatedAt  time.Time `xorm:"created"`
}

func (this *PrizeHistoryModel) TableName() string {
	return "prize_history"
}

//取最后一条领取记录
func (this *PrizeHistoryModel) GetByActivityWuId(activityId, wuid int64) *PrizeHistoryModel {
	history := new(PrizeHistoryModel)
	history.Wuid = wuid
	history.ActivityId = activityId
	has, _ := config.GetDbR(APP_DB_READ).Desc("id").Get(history)
	if has == false {
		return nil
	}
	return history
}

func (this *PrizeHistoryModel) LimitByActivityList(activityId int64, index int, limit int) (histories []*PrizeHistoryModel) {
	err := config.GetDbR(APP_DB_READ).Where("activity_id = ?", activityId).Limit(limit, (index-1)*limit).Find(&histories)
	if err != nil {
		return nil
	}
	return histories
}

func (this *PrizeHistoryModel) Insert(model *PrizeHistoryModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(model)
}

func (this *PrizeHistoryModel) DeleteById(id int64) bool {
	if id == 0 {
		return false
	}
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&PrizeHistoryModel{})
	if err != nil {
		return false
	}
	return true
}

func (this *PrizeHistoryModel) Count(history *PrizeHistoryModel) int64 {
	total, err := config.GetDbW(APP_DB_WRITE).Count(history)
	if err != nil {
		return 0
	}
	return total
}
