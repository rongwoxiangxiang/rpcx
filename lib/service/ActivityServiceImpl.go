package service

import (
	"context"
	"github.com/jinzhu/copier"
	"rpc/lib/models"
)

type ActivityServiceImpl struct{}

func (h ActivityServiceImpl) GetById(ctx context.Context, activity *Activity) (*Activity, error) {
	activityRes := new(Activity)
	ActivityModel := new(models.ActivityModel)
	activityData := ActivityModel.GetById(activity.Id)
	copier.Copy(activityRes, activityData)
	return activityRes, nil
}

func (h ActivityServiceImpl) List(ctx context.Context, query *ActivityQuery) (*Activitylist, error) {
	activityRes := &Activitylist{
		Total:    0,
		Pagesize: 0,
	}
	if query.Wid == 0 {
		return activityRes, nil
	}
	total := (&models.ActivityModel{}).Count(&models.ActivityModel{
		Wid: query.Wid,
	})
	if total <= 0 {
		return activityRes, nil
	}
	activities := (&models.ActivityModel{}).LimitUnderWidList(
		int(query.Index),
		int(query.Limit),
		int(query.Wid),
	)
	activityResData := new([]*Activity)
	err := copier.Copy(activityResData, activities)
	if err != nil {
		return activityRes, nil
	}
	activityRes.Total = total
	activityRes.Pagesize = int64(len(*activityResData))
	activityRes.Activities = *activityResData
	return activityRes, nil
}

func (h ActivityServiceImpl) Insert(ctx context.Context, activityPb *Activity) (*Activity, error) {
	activityModel := new(models.ActivityModel)
	copier.Copy(activityModel, activityPb)
	insertId, err := activityModel.Insert(activityModel)
	activityPb.Id = insertId
	return activityPb, err
}
