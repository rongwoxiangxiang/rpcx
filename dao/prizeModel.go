package dao

import (
	"rpc/common"
	"rpc/config"
	"time"
)

type PrizeInterfaceR interface {
	GetOneUsedPrize(int64, string, int64) *PrizeModel
	LimitByActivityId(int64, int, int) []*PrizeModel
	Count(*PrizeModel) int64
}

type PrizeInterfaceW interface {
	UsePrize(prize *PrizeModel) bool
	Insert(*PrizeModel) (int64, error)
	InsertBatch([]*PrizeModel) (int64, error)
	Update(*PrizeModel) (int64, error)
	DeleteById(int64) bool
}

type PrizeModel struct {
	Id         int64 `xorm:"pk"`
	Wid        int64
	ActivityId int64
	Code       string
	Level      string
	Used       bool //tinyint(1)
	CreatedAt  time.Time
}

func (this *PrizeModel) TableName() string {
	return "prizes"
}

func (this *PrizeModel) GetOneUsedPrize(activityId int64, level string, idGt int64) *PrizeModel {
	if activityId < 1 {
		return nil
	}
	if idGt > 0 {
		config.GetDbR(APP_DB_READ).Where("id >= ?", idGt)
	}
	prize := new(PrizeModel)
	has, err := config.GetDbR(APP_DB_READ).Where("activity_id = ? AND level = '?' AND used = ?", activityId, level, common.NO_VALUE).Get(prize)
	if err != nil {
		config.Logger().Errorf("prize model: GetOneUsedPrize err: %t", err)
	} else if has == false {
		config.Logger().Info("prize model: GetOneUsedPrize empty")
	}
	return prize
}

func (this *PrizeModel) UsePrize(prize *PrizeModel) bool {
	if prize == nil || prize.Id < 0 {
		return false
	}
	prize.Used = true
	_, err := config.GetDbW(APP_DB_WRITE).
		Where("id = ? and used = ?", prize.Id, 0).
		Cols("used").
		Update(prize)
	if err != nil {
		return false
	}
	return true
}

func (this *PrizeModel) Insert(prize *PrizeModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(prize)
}

func (this *PrizeModel) InsertBatch(prizes []*PrizeModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).Insert(&prizes)
}

func (this *PrizeModel) Update(prize *PrizeModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).Id(prize.Id).Update(prize)
}

func (this *PrizeModel) DeleteById(id int64) bool {
	if id == 0 {
		return false
	}
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&PrizeModel{})
	if err != nil {
		return false
	}
	return true
}

func (this *PrizeModel) LimitByActivityId(activityId int64, index int, limit int) (prizes []*PrizeModel) {
	if activityId == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("acitivity_id = ?", activityId).Limit(limit, (index-1)*limit).Find(&prizes)
	if err != nil {
		return nil
	}
	return prizes
}

func (this *PrizeModel) Count(prize *PrizeModel) int64 {
	total, err := config.GetDbW(APP_DB_WRITE).Count(prize)
	if err != nil {
		return 0
	}
	return total
}
