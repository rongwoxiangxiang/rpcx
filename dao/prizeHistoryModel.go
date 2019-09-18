package dao

import (
	"rpc/common"
	"rpc/config"
	"time"
)

type PrizeHistoryInterfaceR interface {
	GetByActivityWuId(activityId, wuid int64) (*PrizeHistoryModel, error)
	LimitUnderActivityList(activityId int64, index int, limit int) []*PrizeHistoryModel
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
	Level      int8
	CreatedAt  time.Time
}

func (this *PrizeHistoryModel) TableName() string {
	return "prize_history"
}

//取最后一条领取记录
func (this *PrizeHistoryModel) GetByActivityWuId(activityId, wuid int64) (*PrizeHistoryModel, error) {
	history := new(PrizeHistoryModel)
	history.Wuid = wuid
	history.ActivityId = activityId
	has, err := config.GetDbR(APP_DB_READ).Desc("id").Get(history)
	if err != nil {
		return nil, common.ErrDataGet
	} else if has == false {
		return nil, common.ErrDataEmpty
	}
	return history, nil
}

func (this *PrizeHistoryModel) LimitUnderActivityList(activityId int64, index int, limit int) (histories []*PrizeHistoryModel) {
	if activityId == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("acitivity_id = ?", activityId).Limit(limit, (index-1)*limit).Find(&histories)
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
