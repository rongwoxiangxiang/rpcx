package dao

import (
	"math/rand"
	"rpc/common"
	"rpc/config"
	"time"
)

type LotteryInterfaceR interface {
	List(wid, activityId int64) []*LotteryModel
}

type LotteryInterfaceW interface {
	Insert(*LotteryModel) (int64, error)
	DeleteById(int64) bool
	Luck(wid, activityId int64) (*LotteryModel, error)
}

type LotteryModel struct {
	Id          int64
	Wid         int64
	ActivityId  int64
	Name        string
	Desc        string
	TotalNum    int64
	ClaimedNum  int64
	Probability int
	FirstCodeId int64
	Level       int8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (this *LotteryModel) TableName() string {
	return "lotteries"
}

func (this *LotteryModel) Insert(model *LotteryModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(model)
}

func (this *LotteryModel) DeleteById(id int64) bool {
	if id == 0 {
		return false
	}
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&LotteryModel{})
	if err != nil {
		return false
	}
	return true
}

//抽奖
func (this *LotteryModel) Luck(wid, activityId int64) (lottery *LotteryModel, err error) {
	lotteries := this.List(wid, activityId)
	if len(lotteries) < 1 {
		return nil, common.ErrDataUnExist
	}
	max := MAX_LUCKY_NUM
	actvityFinished := true
	for _, lot := range lotteries {
		if lot.ClaimedNum >= lot.TotalNum { //当前奖品发放完毕
			max -= lot.Probability
			continue
		}
		actvityFinished = false
		random := rand.Intn(max)
		if random <= lot.Probability {
			lottery = lot
			break
		}
		max -= lot.Probability
	}
	if actvityFinished { //全部奖品发放完毕，自动结束活动
		(&ReplyModel{Wid: this.Wid, ActivityId: this.ActivityId}).ChangeDisabledByWidActivityId(wid, activityId, common.YES_VALUE)
		return nil, common.ErrLuckFinal
	}

	if lottery == nil {
		return nil, common.ErrLuckFail
	}
	claimedNum := lottery.ClaimedNum
	lottery.ClaimedNum = claimedNum + 1
	_, err = config.GetDbW(APP_DB_WRITE).
		Table(new(LotteryModel)).
		Where("id = ? and claimed_num = ?", lottery.Id, claimedNum).
		Cols("claimed_num").
		Update(lottery)
	if err != nil {
		return nil, common.ErrDataUpdate
	}
	return lottery, err
}

func (this *LotteryModel) List(wid, activityId int64) (lotteries []*LotteryModel) {
	if activityId == 0 || wid == 0 {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("wid = ? and activity_id = ?", wid, activityId).OrderBy("probability desc").Find(&lotteries)
	if err != nil {
		lotteries = nil
	}
	return lotteries
}
