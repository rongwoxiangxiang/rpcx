package dao

import (
	"rpc/config"
	"strconv"
	"time"
)

type ReplyInterfaceR interface {
	GetById(int64) *ReplyModel
	FindOneByAliasOrClickKey(int64, string, string) *ReplyModel
	LimitByWid(wid int64, index int, limit int) []*ReplyModel
	Count(*ReplyModel) int64
}

type ReplyInterfaceW interface {
	Insert(*ReplyModel) (int64, error)
	ChangeStatusByWidActivityId(wid, activityId int64, status int8) bool
	Update(*ReplyModel) (int64, error)
	DeleteById(int64) bool
}

type ReplyModel struct {
	Id            int64 `xorm:"pk"`
	Wid           int64
	ActivityId    int64
	Alias         string
	ClickKey      string
	Success       string
	Fail          string //活动数据不存在，未找到等报错是返回信息
	NoPrizeReturn string
	Extra         string
	Type          string
	Status        int8
	Match         string
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
}

func (r *ReplyModel) TableName() string {
	return "replies"
}

func (r *ReplyModel) GetById(id int64) *ReplyModel {
	if id != 0 {
		user := new(ReplyModel)
		config.CacheGetStruct(r.TableName()+strconv.FormatInt(id, 10), user)
		if user != nil && user.Id > 0 {
			return user
		}
		user.Id = id
		has, _ := config.GetDbR(APP_DB_READ).Get(user)
		if has {
			config.CacheSetJson(r.TableName()+strconv.FormatInt(id, 10), user, 3600*24*10)
			return user
		}
	}
	return nil
}

/**
 * @Find
 * @Param Reply.Id int
 * @Param Reply.Alias string
 * @Param Reply.ClickKey string
 * @Success Reply
 */
func (r *ReplyModel) FindOneByAliasOrClickKey(wid int64, alias, clickKey string) *ReplyModel {
	if wid < 0 || ("" == alias && "" == clickKey) {
		return nil
	}
	qs := config.GetDbR(APP_DB_READ).Where("wid = ?", wid)
	if "" != alias {
		qs = qs.Where("alias = ?", alias)
	} else if "" != clickKey {
		qs = qs.Where("click_key = ?", clickKey)
	}
	reply := new(ReplyModel)
	has, err := qs.Where("status = ?", REPLT_STATUS_OPEN).Get(reply)
	if !has || err != nil {
		return nil
	}
	return reply
}

func (r *ReplyModel) LimitByWid(wid int64, index int, limit int) (relpies []*ReplyModel) {
	if wid == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("wid = ?", wid).Limit(limit, (index-1)*limit).Find(&relpies)
	if err != nil {
		return nil
	}
	return relpies
}

func (this *ReplyModel) Count(reply *ReplyModel) int64 {
	total, err := config.GetDbW(APP_DB_WRITE).Count(reply)
	if err != nil {
		return 0
	}
	return total
}

func (r *ReplyModel) ChangeStatusByWidActivityId(wid, activityId int64, status int8) bool {
	if wid < 1 || activityId < 1 {
		return false
	}
	reply := ReplyModel{Wid: wid, ActivityId: activityId}
	has, err := config.GetDbR(APP_DB_READ).Get(&reply)
	if err != nil || has == false {
		return false
	}
	reply.Status = status
	_, err = config.GetDbW(APP_DB_WRITE).Id(reply.Id).Cols("status").Update(reply)
	if err != nil {
		return false
	}
	return true
}

func (r *ReplyModel) Insert(model *ReplyModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(model)
}

func (r *ReplyModel) Update(model *ReplyModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).Id(model.Id).Update(model)
}

func (r *ReplyModel) DeleteById(id int64) bool {
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(new(ReplyModel))
	if err != nil {
		return false
	}
	return true
}
