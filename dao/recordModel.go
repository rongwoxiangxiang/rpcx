package dao

import (
	"rpc/config"
	"time"
)

type RecordInterfaceR interface {
	GetById(int64) *RecordModel
	LimitByWid(int64, int, int) []*RecordModel
	LimitByWuid(int64, int, int) []*RecordModel
	Count(*RecordModel) int64
}

type RecordInterfaceW interface {
	Insert(*RecordModel) (int64, error)
}

type RecordModel struct {
	Id        int64 `xorm:"pk"`
	Wid       int64
	Wuid      int64
	Type      string
	Content   string
	CreatedAt time.Time `xorm:"created"`
}

func (this *RecordModel) TableName() string {
	return "records"
}

func (this *RecordModel) Insert(model *RecordModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(model)
}

func (this *RecordModel) GetById(id int64) *RecordModel {
	if id != 0 {
		record := new(RecordModel)
		record.Id = id
		has, err := config.GetDbR(APP_DB_READ).Get(record)
		if !has || err != nil {
			return nil
		}
		return record
	}

	return nil
}

//通过wid查找
func (r *RecordModel) LimitByWid(wid int64, index int, limit int) (records []*RecordModel) {
	if wid == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("wid = ?", wid).Limit(limit, (index-1)*limit).Find(&records)
	if err != nil {
		return nil
	}
	return records
}

//通过wuid查找
func (r *RecordModel) LimitByWuid(wuid int64, index int, limit int) (records []*RecordModel) {
	if wuid == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("wuid = ?", wuid).Limit(limit, (index-1)*limit).Find(&records)
	if err != nil {
		return nil
	}
	return records
}

func (this *RecordModel) Count(record *RecordModel) int64 {
	total, err := config.GetDbW(APP_DB_WRITE).Count(record)
	if err != nil {
		return 0
	}
	return total
}
