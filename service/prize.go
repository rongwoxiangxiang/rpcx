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
func CopyPrizeHistoryDaoToPb(history *dao.PrizeHistoryModel, pbHistory *pb.PrizeHistory) *pb.PrizeHistory {
	if pbHistory == nil {
		pbHistory = new(pb.PrizeHistory)
	}
	pbHistory.Id = history.Id
	pbHistory.Wuid = history.Wuid
	pbHistory.ActivityId = history.ActivityId
	pbHistory.Code = history.Code
	pbHistory.Level = history.Level
	pbHistory.Prize = history.Prize
	pbHistory.CreatedAt = history.CreatedAt.Unix()
	return pbHistory
}

func (this *PrizeService) LimitPrizeByActivityIdAndUsed(ctx context.Context, resq *pb.RequestList, resp *pb.PrizeList) error {
	var (
		wechatId int64
		usedInt8 int8
	)
	if wid, ok := resq.Params["wid"]; ok {
		wechatId = util.StringToInt64(wid)
	}
	if used, ok := resq.Params["used"]; ok { //若传值不为1默认为2
		usedInt8 = common.NO_VALUE
		if used == common.YES_VALUE_STRING {
			usedInt8 = common.YES_VALUE
		}
	}
	prizes := dao.GetPrizeServiceR().LimitByActivityIdAndUsed(
		wechatId, usedInt8, *(*int)(unsafe.Pointer(&resq.Index)), *(*int)(unsafe.Pointer(&resq.Limit)))
	for _, prize := range prizes {
		resp.Prizes = append(resp.Prizes, CopyPrizeDaoToPb(prize, nil))
	}
	prize := &dao.PrizeModel{Wid: wechatId}
	prize.SetUsed(usedInt8)
	resp.Total = dao.GetPrizeServiceR().Count(prize)
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	return nil
}

func (this *PrizeService) InsertPrize(ctx context.Context, resq *pb.Prize, resp *pb.Prize) error {
	prize := &dao.PrizeModel{
		Wid:        resq.Wid,
		ActivityId: resq.ActivityId,
		Code:       resq.Code,
		Level:      resq.Level,
		Used:       common.NO_VALUE,
		CreatedAt:  time.Unix(resq.CreatedAt, 0),
	}
	_, err := dao.GetPrizeServiceW().Insert(prize)
	if err == nil {
		CopyPrizeDaoToPb(prize, resp)
	}
	return err
}

func (this *PrizeService) DeletePrize(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
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
	if prize == nil || prize.Id < 1 {
		return common.ErrLuckFinal
	}
	ok := dao.GetPrizeServiceW().UsePrize(prize)
	if !ok {
		config.Logger().Infof("Prize service:use prize err: %v", prize)
	}
	CopyPrizeDaoToPb(prize, resp)
	return nil
}

func (this *PrizeService) BatchInsertPrize(ctx context.Context, resq *pb.PrizeAdd, resp *pb.ResponseEffect) error {
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

func (this *PrizeService) GetByActivityWuId(ctx context.Context, resq *pb.RequestOneByActivityWuId, resp *pb.PrizeHistory) error {
	history, err := dao.GetPrizeHistoryServiceR().GetByActivityWuId(resq.ActivityId, resq.Wuid)
	if err != nil {
		return err
	}
	CopyPrizeHistoryDaoToPb(history, resp)
	return nil
}

func (this *PrizeService) LimitByActivityList(ctx context.Context, resq *pb.RequestList, resp *pb.PrizeHistoryList) error {
	activityId, ok := resq.Params["activityId"]
	if !ok {
		return nil
	}
	activityIdInt64 := util.StringToInt64(activityId)
	histories := dao.GetPrizeHistoryServiceR().LimitByActivityList(
		activityIdInt64, *(*int)(unsafe.Pointer(&resq.Index)), *(*int)(unsafe.Pointer(&resq.Limit)))
	for _, history := range histories {
		resp.Histories = append(resp.Histories, CopyPrizeHistoryDaoToPb(history, nil))
	}
	resp.Total = dao.GetPrizeHistoryServiceR().Count(&dao.PrizeHistoryModel{ActivityId: activityIdInt64})
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	return nil
}

func (this *PrizeService) InsertPrizeHistory(ctx context.Context, resq *pb.PrizeHistory, resp *pb.PrizeHistory) error {
	history := &dao.PrizeHistoryModel{
		Wuid:       resq.Wuid,
		ActivityId: resq.ActivityId,
		Code:       resq.Code,
		Level:      resq.Level,
		Prize:      resq.Prize,
		CreatedAt:  time.Unix(resq.CreatedAt, 0),
	}
	_, err := dao.GetPrizeHistoryServiceW().Insert(history)
	if err == nil {
		CopyPrizeHistoryDaoToPb(history, resp)
	}
	return err
}

func (this *PrizeService) DeletePrizeHistory(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	ok := dao.GetPrizeHistoryServiceW().DeleteById(resq.Id)
	if ok {
		resp.Success = true
		resp.Effect = 1
	}
	return nil
}
