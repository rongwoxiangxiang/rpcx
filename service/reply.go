package service

import (
	"context"
	"rpc/common"
	"rpc/dao"
	"rpc/pb"
	"rpc/util"
	"time"
	"unsafe"
)

type ReplyService struct {
}

func CopyReplyDaoToPb(record *dao.ReplyModel, pbReply *pb.Reply) *pb.Reply {
	if pbReply == nil {
		pbReply = new(pb.Reply)
	}
	pbReply.Id = record.Id
	pbReply.Wid = record.Wid
	pbReply.ActivityId = record.ActivityId
	pbReply.Alias = record.Alias
	pbReply.ClickKey = record.ClickKey
	pbReply.Success = record.Success
	pbReply.Fail = record.Fail
	pbReply.NoPrizeReturn = record.NoPrizeReturn
	pbReply.Extra = record.Extra
	pbReply.Status = pb.ReplyStatus(record.Status)
	pbReply.Match = record.Match
	pbReply.CreatedAt = record.CreatedAt.Unix()
	pbReply.UpdatedAt = record.UpdatedAt.Unix()
	return pbReply
}

func (this *ReplyService) GetById(ctx context.Context, resq *pb.RequestById, resp *pb.Reply) error {
	if resq.Id < 1 {
		return nil
	}
	reply := dao.GetReplyServiceR().GetById(resq.Id)
	CopyReplyDaoToPb(reply, resp)
	return nil
}

func (this *ReplyService) FindOneByAliasOrClickKey(ctx context.Context, resq *pb.Reply, resp *pb.Reply) error {
	reply := dao.GetReplyServiceR().FindOneByAliasOrClickKey(resq.Wid, resq.Alias, resq.ClickKey)
	CopyReplyDaoToPb(reply, resp)
	return nil
}

func (this *ReplyService) LimitByWid(ctx context.Context, resq *pb.RequestList, resp *pb.ReplyList) error {
	wid, ok := resq.Params["wid"]
	if !ok {
		return nil
	}
	widInt64 := util.StringToInt64(wid)
	records := dao.GetReplyServiceR().LimitByWid(
		widInt64,
		*(*int)(unsafe.Pointer(&resq.Index)),
		*(*int)(unsafe.Pointer(&resq.Limit)))
	for _, reply := range records {
		resp.Replies = append(resp.Replies, CopyReplyDaoToPb(reply, nil))
	}
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	resp.Total = dao.GetReplyServiceR().Count(&dao.ReplyModel{Wid: widInt64})
	return nil
}

func (this *ReplyService) Insert(ctx context.Context, resq *pb.Reply, resp *pb.Reply) error {
	record := &dao.ReplyModel{
		Wid:           resq.Wid,
		ActivityId:    resq.ActivityId,
		Alias:         resq.Alias,
		ClickKey:      resq.ClickKey,
		Success:       resq.Success,
		Fail:          resq.Fail,
		NoPrizeReturn: resq.NoPrizeReturn,
		Extra:         resq.Extra,
		Match:         resq.Match,
		Status:        *(*int8)(unsafe.Pointer(&resq.Status)),
		CreatedAt:     time.Unix(resq.CreatedAt, 0),
		UpdatedAt:     time.Unix(resq.UpdatedAt, 0),
	}
	_, err := dao.GetReplyServiceW().Insert(record)
	if err == nil {
		CopyReplyDaoToPb(record, resp)
	}
	return err
}

func (this *ReplyService) ChangeStatusByWidActivityId(ctx context.Context, resq *pb.StatusByWidActivityId, resp *pb.Reply) error {
	ok := dao.GetReplyServiceW().ChangeStatusByWidActivityId(
		resq.Wid,
		resq.ActivityId,
		*(*int8)(unsafe.Pointer(&resq.Status)),
	)
	if !ok {
		return common.ErrDataUpdate
	}
	return nil
}

func (this *ReplyService) Update(ctx context.Context, resq *pb.Reply, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	rows, err := dao.GetReplyServiceW().Update(&dao.ReplyModel{
		Id:            resq.Id,
		Wid:           resq.Wid,
		ActivityId:    resq.ActivityId,
		Alias:         resq.Alias,
		ClickKey:      resq.ClickKey,
		Success:       resq.Success,
		Fail:          resq.Fail,
		NoPrizeReturn: resq.NoPrizeReturn,
		Extra:         resq.Extra,
		Match:         resq.Match,
		Status:        *(*int8)(unsafe.Pointer(&resq.Status)),
		CreatedAt:     time.Unix(resq.CreatedAt, 0),
		UpdatedAt:     time.Unix(resq.UpdatedAt, 0),
	})
	if err == nil {
		resp.Success = true
		resp.Effect = rows
	}
	return err
}

func (this *ReplyService) Delete(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	ok := dao.GetReplyServiceW().DeleteById(resq.Id)
	if ok {
		resp.Success = true
		resp.Effect = 1
	}
	return nil
}
