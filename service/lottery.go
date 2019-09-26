package service

import (
	"context"
	"math/rand"
	"rpc/common"
	"rpc/dao"
	"rpc/pb"
	"time"
)

type LotteryService struct {
}

func CopyLotteryDaoToPb(lottery *dao.LotteryModel, pbLottery *pb.Lottery) *pb.Lottery {
	if pbLottery == nil {
		pbLottery = new(pb.Lottery)
	}
	pbLottery.Id = lottery.Id
	pbLottery.Wid = lottery.Wid
	pbLottery.ActivityId = lottery.ActivityId
	pbLottery.Name = lottery.Name
	pbLottery.Desc = lottery.Desc
	pbLottery.TotalNum = lottery.TotalNum
	pbLottery.ClaimedNum = lottery.ClaimedNum
	pbLottery.Probability = lottery.Probability
	pbLottery.FirstCodeId = lottery.FirstCodeId
	pbLottery.Level = lottery.Level
	pbLottery.CreatedAt = lottery.CreatedAt.Unix()
	pbLottery.UpdatedAt = lottery.UpdatedAt.Unix()
	return pbLottery
}

func (this *LotteryService) ListByWidAndActivityId(ctx context.Context, resq *pb.RequestWidAndActivityId, resp *pb.LotteryList) error {
	lotteries := dao.GetLotteryServiceR().ListByWidAndActivityId(resq.Wid, resq.ActivityId)
	if lotteries == nil {
		return nil
	}
	resp.Total = int64(len(lotteries))
	for _, lottery := range lotteries {
		resp.Lotteries = append(resp.Lotteries, CopyLotteryDaoToPb(lottery, nil))
	}
	return nil
}

func (this *LotteryService) Luck(ctx context.Context, resq *pb.RequestWidAndActivityId, resp *pb.Lottery) (err error) {
	lotteries := dao.GetLotteryServiceR().ListByWidAndActivityId(resq.Wid, resq.ActivityId)
	if lotteries == nil {
		return common.ErrDataUnExist
	}
	max := int32(dao.MAX_LUCKY_NUM)
	actvityFinished := true
	lottery := new(dao.LotteryModel)
	for _, lot := range lotteries {
		if lot.ClaimedNum >= lot.TotalNum { //当前奖品发放完毕
			max -= lot.Probability
			continue
		}
		actvityFinished = false
		random := rand.Int31n(max)
		if random <= lot.Probability { //抽到奖品
			lottery = lot
			break
		}
		max -= lot.Probability
	}
	if actvityFinished { //全部奖品发放完毕，自动关闭回复
		dao.GetReplyServiceW().ChangeStatusByWidActivityId(resq.Wid, resq.ActivityId, dao.REPLT_STATUS_MANUAL_CLOSE)

		return common.ErrLuckFinal
	}
	if lottery == nil || lottery.Id < 1 {
		return common.ErrLuckFail
	}
	err = dao.GetLotteryServiceW().IncrClaimedNum(lottery)
	if err != nil {
		return common.ErrDataUpdate
	}
	return
}

func (this *LotteryService) Insert(ctx context.Context, resq *pb.Lottery, resp *pb.Lottery) error {
	lottery := &dao.LotteryModel{
		Name:        resq.Name,
		Wid:         resq.Wid,
		ActivityId:  resq.ActivityId,
		Desc:        resq.Desc,
		TotalNum:    resq.TotalNum,
		ClaimedNum:  resq.ClaimedNum,
		Probability: resq.Probability,
		FirstCodeId: resq.FirstCodeId,
		Level:       resq.Level,
		CreatedAt:   time.Unix(resq.CreatedAt, 0),
		UpdatedAt:   time.Unix(resq.UpdatedAt, 0),
	}
	_, err := dao.GetLotteryServiceW().Insert(lottery)
	if err == nil {
		CopyLotteryDaoToPb(lottery, resp)
	}
	return err
}

func (this *LotteryService) Delete(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	ok := dao.GetLotteryServiceW().DeleteById(resq.Id)
	if ok {
		resp.Success = true
		resp.Effect = 1
	}
	return nil
}
