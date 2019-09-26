package service

import (
	"context"
	"rpc/dao"
	"rpc/pb"
	"rpc/util"
	"time"
	"unsafe"
)

type ActivityService struct {
}

func CopyActivityDaoToPb(activity *dao.ActivityModel, pbActivity *pb.Activity) *pb.Activity {
	if pbActivity == nil {
		pbActivity = new(pb.Activity)
	}
	pbActivity.Id = activity.Id
	pbActivity.Wid = activity.Wid
	pbActivity.Name = activity.Name
	pbActivity.Desc = activity.Desc
	pbActivity.ActivityType = activity.ActivityType
	pbActivity.RelativeId = activity.RelativeId
	pbActivity.TimeStarted = activity.TimeStarted.Unix()
	pbActivity.TimeEnd = activity.TimeEnd.Unix()
	pbActivity.CreatedAt = activity.CreatedAt.Unix()
	pbActivity.UpdatedAt = activity.UpdatedAt.Unix()
	return pbActivity
}

func (this *ActivityService) GetById(ctx context.Context, resq *pb.RequestById, resp *pb.Activity) error {
	if resq.Id < 1 {
		return nil
	}
	activity := dao.GetActivityServiceR().GetById(resq.Id)
	CopyActivityDaoToPb(activity, resp)
	return nil
}

func (this *ActivityService) LimitByWid(ctx context.Context, resq *pb.RequestList, resp *pb.ActivityList) error {
	wid, ok := resq.Params["wid"]
	if !ok {
		return nil
	}
	widInt64 := util.StringToInt64(wid)
	activities := dao.GetActivityServiceR().ListByWid(
		widInt64,
		*(*int)(unsafe.Pointer(&resq.Index)),
		*(*int)(unsafe.Pointer(&resq.Limit)))
	for _, activity := range activities {
		resp.Activities = append(resp.Activities, CopyActivityDaoToPb(activity, nil))
	}
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	resp.Total = dao.GetActivityServiceR().Count(&dao.ActivityModel{Wid: widInt64})
	return nil
}

func (this *ActivityService) Insert(ctx context.Context, resq *pb.Activity, resp *pb.Activity) error {
	activity := &dao.ActivityModel{
		Name:         resq.Name,
		Desc:         resq.Desc,
		ActivityType: resq.ActivityType,
		RelativeId:   resq.RelativeId,
		TimeStarted:  time.Unix(resq.TimeStarted, 0),
		TimeEnd:      time.Unix(resq.TimeEnd, 0),
	}
	_, err := dao.GetActivityServiceW().Insert(activity)
	if err == nil {
		CopyActivityDaoToPb(activity, resp)
	}
	return err
}

func (this *ActivityService) Update(ctx context.Context, resq *pb.Activity, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	rows, err := dao.GetActivityServiceW().Update(&dao.ActivityModel{
		Id:           resq.Id,
		Name:         resq.Name,
		Desc:         resq.Desc,
		ActivityType: resq.ActivityType,
		RelativeId:   resq.RelativeId,
		TimeStarted:  time.Unix(resq.TimeStarted, 0),
		TimeEnd:      time.Unix(resq.TimeEnd, 0),
	})
	if err == nil {
		resp.Success = true
		resp.Effect = rows
	}
	return err
}

func (this *ActivityService) Delete(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	ok := dao.GetActivityServiceW().DeleteById(resq.Id)
	if ok {
		resp.Success = true
		resp.Effect = 1
	}
	return nil
}
