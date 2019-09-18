package service

import (
	"context"
	"rpc/dao"
	"rpc/pb"
	"rpc/util"
	"time"
)

type CheckinService struct {
}

func CopyCheckinDaoToPb(checkin *dao.CheckinModel, pbCheckin *pb.Checkin) *pb.Checkin {
	if pbCheckin == nil {
		pbCheckin = new(pb.Checkin)
	}
	pbCheckin.Id = checkin.Id
	pbCheckin.Wid = checkin.Wid
	pbCheckin.Total = checkin.Total
	pbCheckin.ActivityId = checkin.ActivityId
	pbCheckin.Wuid = checkin.Wuid
	pbCheckin.Liner = checkin.Liner
	pbCheckin.Lastcheckin = checkin.Lastcheckin.Unix()
	pbCheckin.CreatedAt = checkin.CreatedAt.Unix()
	pbCheckin.UpdatedAt = checkin.UpdatedAt.Unix()
	return pbCheckin
}

func (this *CheckinService) GetCheckinInfoByActivityIdAndWuid(ctx context.Context, resq *pb.Checkin, resp *pb.Checkin) error {
	checkin, err := dao.GetCheckinServiceR().GetCheckinInfoByActivityIdAndWuid(resq.ActivityId, resp.Wuid)
	if err != nil {
		return err
	}
	if checkin == nil {
		checkin = &dao.CheckinModel{
			ActivityId: resp.ActivityId,
			Wuid:       resp.Wuid,
		}
		_, err = dao.GetCheckinServiceW().Insert(checkin)
	}
	CopyCheckinDaoToPb(checkin, resp)
	return err
}

func (this *CheckinService) LimitByWid(ctx context.Context, resq *pb.RequestList, resp *pb.CheckinList) error {
	wid, ok := resq.Params["wid"]
	if !ok {
		return nil
	}
	widInt64 := util.StringToInt64(wid)
	checkins := dao.GetCheckinServiceR().ListByWid(widInt64)
	for _, checkin := range checkins {
		resp.Checkins = append(resp.Checkins, CopyCheckinDaoToPb(checkin, nil))
	}
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	resp.Total = dao.GetCheckinServiceR().Count(&dao.CheckinModel{Wid: widInt64})
	return nil
}

func (this *CheckinService) Insert(ctx context.Context, resq *pb.Checkin, resp *pb.Checkin) error {
	checkin := &dao.CheckinModel{
		Wid:         resq.Wid,
		Total:       resq.Total,
		ActivityId:  resq.ActivityId,
		Wuid:        resq.Wuid,
		Liner:       resq.Liner,
		Lastcheckin: time.Unix(resq.Lastcheckin, 0),
		CreatedAt:   time.Unix(resq.CreatedAt, 0),
		UpdatedAt:   time.Unix(resq.UpdatedAt, 0),
	}
	_, err := dao.GetCheckinServiceW().Insert(checkin)
	if err == nil {
		CopyCheckinDaoToPb(checkin, resp)
	}
	return err
}

func (this *CheckinService) Update(ctx context.Context, resq *pb.Checkin, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	rows, err := dao.GetCheckinServiceW().Update(&dao.CheckinModel{
		Id:          resq.Id,
		Total:       resq.Total,
		ActivityId:  resq.ActivityId,
		Liner:       resq.Liner,
		Lastcheckin: time.Unix(resq.Lastcheckin, 0),
		CreatedAt:   time.Unix(resq.CreatedAt, 0),
	})
	if err == nil {
		resp.Success = true
		resp.Effect = rows
	}
	return nil
}

func (this *CheckinService) Delete(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	ok := dao.GetCheckinServiceW().DeleteById(resq.Id)
	if ok {
		resp.Success = true
		resp.Effect = 1
	}
	return nil
}
