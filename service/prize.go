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
	pbPrize.Used = Prize.GetBoolUsed()
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
		CreatedAt:  time.Unix(resq.CreatedAt, 0),
	}
	Prize.SetUsed(resq.Used)
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

func (this *PrizeService) InsertBatch(ctx context.Context, resq *pb.PrizeAdd, resp *pb.ResponseEffect) error {
	wechat := dao.GetWechatServiceR().GetById(resq.Wid)
	activity := dao.GetActivityServiceR().GetById(resq.ActivityId)
	if activity == nil || wechat == nil {
		config.Logger().Infof("Prize service batch insert activity[%d] or wechat[%d] not exist", resq.ActivityId, resq.Wid)
		return nil
	}
	prizes := make([]*dao.PrizeModel, 0)
	resp.Success = true
	for _, prize := range resq.Codes {
		prizes = append(prizes, &dao.PrizeModel{
			Wid: resq.Wid, Level: resq.Level, ActivityId: resq.ActivityId, Code: prize, Used: common.NO_VALUE,
		})
		if len(prizes) == dao.INSER_DEFAULT_ROWS_EACH { //每INSER_DEFAULT_ROWS_EACH insert次
			rows, err := dao.GetPrizeServiceW().InsertBatch(prizes)
			if err != nil {
				config.Logger().Errorf("Prize service bash insert err: %v", err)
			}
			prizes = (prizes)[0:0] //清空
			resp.Effect += rows
		}
	}
	if prizes != nil && len(prizes) > 0 {
		rows, err := dao.GetPrizeServiceW().InsertBatch(prizes)
		if err != nil {
			config.Logger().Errorf("Prize service bash insert err: %v", err)
		}
		resp.Effect += rows
	}
	return nil
}
