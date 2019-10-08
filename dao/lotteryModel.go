package dao

import (
	"rpc/config"
	"time"
)

type LotteryInterfaceR interface {
	ListByWidAndActivityId(int64, int64) []*LotteryModel
	Count(*LotteryModel) int64
}

type LotteryInterfaceW interface {
	Insert(*LotteryModel) (int64, error)
	DeleteById(int64) bool
	IncrClaimedNum(*LotteryModel) error
}

type LotteryModel struct {
	Id          int64
	Wid         int64
	ActivityId  int64
	Name        string
	Desc        string
	TotalNum    int64
	ClaimedNum  int64
	Probability int32
	FirstCodeId int64
	Level       string
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}

func (this *LotteryModel) TableName() string {
	return "lotteries"
}

func (this *LotteryModel) Insert(model *LotteryModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(model)
}

func (this *LotteryModel) DeleteById(id int64) bool {
	if id < 1 {
		return false
	}
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&LotteryModel{})
	if err != nil {
		config.Logger().WithField("id", id).Errorf("LotteryModel DeleteById err: %v", err)
		return false
	}
	return true
}

func (this *LotteryModel) IncrClaimedNum(lottery *LotteryModel) error {
	claimedNum := this.ClaimedNum + 1
	_, err := config.GetDbW(APP_DB_WRITE).
		Table(new(LotteryModel)).
		Where("id = ? and claimed_num = ?", this.Id, claimedNum).
		Cols("claimed_num").
		Update(lottery)
	if err != nil {
		config.Logger().Errorf("LotteryModel IncrClaimedNum Lottery:%v, err: %v", this, err)
	}
	return err
}

func (this *LotteryModel) ListByWidAndActivityId(wid, activityId int64) (lotteries []*LotteryModel) {
	err := config.GetDbR(APP_DB_READ).
		Where("wid = ? and activity_id = ?", wid, activityId).
		OrderBy("probability desc").
		Find(&lotteries)
	if err != nil {
		lotteries = nil
	}
	return lotteries
}

func (this *LotteryModel) Count(lottery *LotteryModel) int64 {
	total, err := config.GetDbW(APP_DB_WRITE).Count(lottery)
	if err != nil {
		return 0
	}
	return total
}
