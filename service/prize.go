package service

import (
	"context"
	"rpc/common"
	"rpc/config"
	"rpc/dao"
	"rpc/pb"
	"rpc/util"
	"time"
	"unsafe"
)

type PrizeService struct {
}

func CopyPrizeDaoToPb(Prize *dao.PrizeModel, pbPrize *pb.Prize) *pb.Prize {
	if pbPrize == nil {
		pbPrize = new(pb.Prize)
	}
	pbPrize.Id = Prize.Id
	pbPrize.Wid = Prize.Wid
	pbPrize.ActivityId = Prize.ActivityId
	pbPrize.Code = Prize.Code
	pbPrize.Level = Prize.Level
	pbPrize.Used = Prize.Used
	pbPrize.CreatedAt = Prize.CreatedAt.Unix()
	return pbPrize
}

func (this *PrizeService) LimitByActivityId(ctx context.Context, resq *pb.RequestList, resp *pb.PrizeList) error {
	wid, ok := resq.Params["wid"]
	if !ok {
		return nil
	}
	widInt64 := util.StringToInt64(wid)
	Prizes := dao.GetPrizeServiceR().LimitByActivityId(
		widInt64,
		*(*int)(unsafe.Pointer(&resq.Index)),
		*(*int)(unsafe.Pointer(&resq.Limit)))
	for _, Prize := range Prizes {
		resp.Prizes = append(resp.Prizes, CopyPrizeDaoToPb(Prize, nil))
	}
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	resp.Total = dao.GetPrizeServiceR().Count(&dao.PrizeModel{Wid: widInt64})
	return nil
}

func (this *PrizeService) Insert(ctx context.Context, resq *pb.Prize, resp *pb.Prize) error {
	Prize := &dao.PrizeModel{
		Wid:        resq.Wid,
		ActivityId: resq.ActivityId,
		Code:       resq.Code,
		Level:      resq.Level,
		Used:       resq.Used,
		CreatedAt:  time.Unix(resq.CreatedAt, 0),
	}
	_, err := dao.GetPrizeServiceW().Insert(Prize)
	if err == nil {
		CopyPrizeDaoToPb(Prize, resp)
	}
	return err
}

func (this *PrizeService) Delete(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	ok := dao.GetPrizeServiceW().DeleteById(resq.Id)
	if ok {
		resp.Success = true
		resp.Effect = 1
	}
	return nil
}

func (this *PrizeService) ChooseOneUsedPrize(ctx context.Context, resq *pb.ChooseOneUsedPrize, resp *pb.Prize) error {
	prize := dao.GetPrizeServiceR().GetOneUsedPrize(resq.ActivityId, resq.Level, resq.GetIdGt())
	if prize == nil {
		return common.ErrLuckFinal
	}
	ok := dao.GetPrizeServiceW().UsePrize(prize)
	if !ok {
		config.Logger().Infof("Prize service:use prize err: %v", prize)
	}
	CopyPrizeDaoToPb(prize, resp)
	return nil
}
