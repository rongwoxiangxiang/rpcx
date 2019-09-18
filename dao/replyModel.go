package dao

import (
	"rpc/common"
	"rpc/config"
	"strconv"
	"time"
)

type ReplyInterfaceR interface {
	GetById(int64) *ReplyModel
	FindOne(*ReplyModel) *ReplyModel
	LimitUnderWidList(wid int64, index int, limit int) []*ReplyModel
}

type ReplyInterfaceW interface {
	Insert(*ReplyModel) (int64, error)
	ChangeDisabledByWidActivityId(wid, activityId int64, disabled int8) bool
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
	Disabled      int8
	Match         int8
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
func (r *ReplyModel) FindOne(model *ReplyModel) *ReplyModel {
	if model.Wid == 0 || ("" == model.Alias && "" == model.ClickKey) {
		return nil
	}
	qs := config.GetDbR(APP_DB_READ).Where("wid = ?", model.Wid)
	if "" != model.Alias {
		qs = qs.Where("alias = ?", model.Alias)
	} else if "" != r.ClickKey {
		qs = qs.Where("click_key = ?", model.ClickKey)
	}
	reply := new(ReplyModel)
	has, err := qs.Where("disabled = ?", common.NO_VALUE).Get(reply)
	if !has || err != nil {
		return nil
	}
	return reply
}

func (r *ReplyModel) LimitUnderWidList(wid int64, index int, limit int) (relpies []*ReplyModel) {
	if wid == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("wid = ?", wid).Limit(limit, (index-1)*limit).Find(&relpies)
	if err != nil {
		return nil
	}
	return relpies
}

func (r *ReplyModel) ChangeDisabledByWidActivityId(wid, activityId int64, disabled int8) bool {
	if wid == 0 || activityId == 0 {
		return false
	}
	reply := ReplyModel{Wid: wid, ActivityId: activityId}
	has, err := config.GetDbR(APP_DB_READ).Get(&reply)
	if err != nil || has == false {
		return false
	}
	reply.Disabled = disabled
	_, err = config.GetDbW(APP_DB_WRITE).Id(reply.Id).Cols("disabled").Update(reply)
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
