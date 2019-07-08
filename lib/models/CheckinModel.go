package models

import (
	"chs/common"
	"rpc/lib/config"
	"time"
)

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
	err := config.GetDefaultR().Where("wid = ?", wid).Find(&lotteries)
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
func (this *CheckinModel) GetCheckinInfoByActivityIdAndWuid(activityId, wuid int64) (checkin *CheckinModel, err error) {
	if activityId == 0 || wuid == 0 {
		err = common.ErrDataGet
		return
	}
	checkin.ActivityId = activityId
	checkin.Wuid = wuid
	has, err := config.GetDefaultR().Get(checkin)
	if !has || err != nil {
		return nil, err
	}
	return
}

/**
 * @GetCheckinByActivityWuid
 * @Description 活动用户签到信息，不存在自动创建
 * @Param id ActivityId
 * @Param id Wuid
 * @Param id Wid
 * @return CheckinModel,error
 */
func (this *CheckinModel) GetCheckinByActivityWuid(activityId, wuid int64) (checkin *CheckinModel, err error) {
	if activityId == 0 || wuid == 0 {
		return nil, common.ErrDataGet
	}
	checkin.ActivityId = activityId
	checkin.Wuid = wuid
	has, err := config.GetDefaultR().Get(&checkin)
	if err != nil {
		return nil, common.ErrDataGet
	} else if !has {
		return nil, common.ErrDataNoExist
	}
	return
}

func (this *CheckinModel) Insert(checkin *CheckinModel) (int64, error) {
	return config.GetDefaultW().InsertOne(checkin)
}

func (this *CheckinModel) Update(checkin *CheckinModel) (int64, error) {
	return config.GetDefaultW().Id(checkin.Id).Update(checkin)
}

func (this *CheckinModel) DeleteById(id int64) bool {
	_, err := config.GetDefaultW().Id(id).Unscoped().Delete(&CheckinModel{})
	if err != nil {
		return false
	}
	return true
}
