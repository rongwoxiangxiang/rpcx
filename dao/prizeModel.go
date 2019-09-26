package dao

import (
	"rpc/common"
	"rpc/config"
	"time"
)

type PrizeInterfaceR interface {
	GetOneUsedPrize(int64, string, int64) *PrizeModel
	LimitByActivityIdAndUsed(int64, int8, int, int) []*PrizeModel
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
	Used       int8
	CreatedAt  time.Time `xorm:"created"`
}

func (this *PrizeModel) TableName() string {
	return "prizes"
}

func (this *PrizeModel) GetBoolUsed() bool {
	if this.Used == common.YES_VALUE {
		return true
	}
	return false
}

func (this *PrizeModel) SetUsed(used int8) {
	if used == common.YES_VALUE {
		this.Used = common.YES_VALUE
	} else if used == common.NO_VALUE {
		this.Used = common.NO_VALUE
	} else {
		config.Logger().Warnf("prize model: SetUsedString with unexpect value: ", used)
	}
}

func (this *PrizeModel) GetOneUsedPrize(activityId int64, level string, idGt int64) *PrizeModel {
	if activityId < 1 {
		return nil
	}
	qs := config.GetDbR(APP_DB_READ).Where("activity_id = ? AND used = ?", activityId, common.NO_VALUE)
	if idGt > 0 {
		qs = qs.Where("id >= ?", idGt)
	}
	if level != "" {
		qs = qs.Where("level = ?", level)
	}
	prize := new(PrizeModel)
	has, err := qs.Get(prize)
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
	prize.Used = common.YES_VALUE
	_, err := config.GetDbW(APP_DB_WRITE).
		Where("id = ? and used = ?", prize.Id, common.NO_VALUE).
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
	//i := 0
	//for lens := len(prizes); ; i++ {
	//	if lens >= INSER_DEFAULT_ROWS_EACH*(i+1) {
	//		prizePre := prizes[i*INSER_DEFAULT_ROWS_EACH : (i+1)*INSER_DEFAULT_ROWS_EACH-1]
	//		_, err = config.GetDbW(APP_DB_WRITE).Insert(&prizePre)
	//		if err != nil {
	//			config.Logger().Errorf("Prize insert batch err: %v", err)
	//		}
	//		if lens == INSER_DEFAULT_ROWS_EACH*(i+1) {
	//			break
	//		}
	//	} else if lens > INSER_DEFAULT_ROWS_EACH*i {
	//		prizePre := prizes[i*INSER_DEFAULT_ROWS_EACH : (i+1)*INSER_DEFAULT_ROWS_EACH]
	//		rows, err = config.GetDbW(APP_DB_WRITE).Insert(&prizePre)
	//		if err != nil {
	//			config.Logger().Errorf("Prize insert batch err: %v", err)
	//		}
	//		break
	//	}
	//}
	//rows += INSER_DEFAULT_ROWS_EACH * int64(i)
	//return rows, err
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

func (this *PrizeModel) LimitByActivityIdAndUsed(activityId int64, used int8, index int, limit int) (prizes []*PrizeModel) {
	if activityId == 0 || index < 0 || limit < 1 {
		return nil
	}
	qs := config.GetDbR(APP_DB_READ).Where("activity_id = ?", activityId).Limit(limit, (index-1)*limit)
	if used == common.YES_VALUE || used == common.NO_VALUE {
		qs = qs.Where("used = ?", used)
	}
	err := qs.Find(&prizes)
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
