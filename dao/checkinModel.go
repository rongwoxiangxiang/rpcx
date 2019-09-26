package dao

import (
	"rpc/common"
	"rpc/config"
	"time"
)

type CheckinInterfaceR interface {
	ListByWid(wid int64) []*CheckinModel
	GetCheckinInfoByActivityIdAndWuid(activityId, wuid int64) (*CheckinModel, error)
	Count(*CheckinModel) int64
}

type CheckinInterfaceW interface {
	Insert(checkin *CheckinModel) (int64, error)
	Update(checkin *CheckinModel) (int64, error)
	DeleteById(id int64) bool
}

type CheckinModel struct {
	Id          int64 `xorm:"pk"`
	Wid         int64
	ActivityId  int64
	Wuid        int64
	Liner       int64
	Total       int64
	Lastcheckin time.Time
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}

func (this *CheckinModel) TableName() string {
	return "checkins"
}

func (this *CheckinModel) ListByWid(wid int64) (lotteries []*CheckinModel) {
	if wid == 0 {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("wid = ?", wid).Find(&lotteries)
	if err != nil {
		return nil
	}
	return lotteries
}

/**
 * @GetCheckinByActivityWuid
 * @Description 活动用户签到信息
 * @Param id ActivityId
 * @Param id Wuid
 * @return CheckinModel,error
 */
func (this *CheckinModel) GetCheckinInfoByActivityIdAndWuid(activityId, wuid int64) (*CheckinModel, error) {
	if activityId == 0 || wuid == 0 {
		return nil, common.ErrDataGet
	}
	checkin := new(CheckinModel)
	checkin.ActivityId = activityId
	checkin.Wuid = wuid
	has, err := config.GetDbR(APP_DB_READ).Get(checkin)
	if !has || err != nil {
		return nil, err
	}
	return checkin, nil
}

func (this *CheckinModel) Count(checkin *CheckinModel) int64 {
	total, err := config.GetDbW(APP_DB_WRITE).Count(checkin)
	if err != nil {
		return 0
	}
	return total
}

func (this *CheckinModel) Insert(checkin *CheckinModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(checkin)
}

func (this *CheckinModel) Update(checkin *CheckinModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).Id(checkin.Id).Update(checkin)
}

func (this *CheckinModel) DeleteById(id int64) bool {
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&CheckinModel{})
	if err != nil {
		return false
	}
	return true
}
