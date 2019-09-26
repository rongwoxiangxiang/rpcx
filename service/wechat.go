package service

import (
	"context"
	"rpc/dao"
	"rpc/pb"
	"rpc/util"
	"time"
)

type WechatService struct {
}

func CopyWechatDaoToPb(wechat *dao.WechatModel, pbWechat *pb.Wechat) *pb.Wechat {
	if pbWechat == nil {
		pbWechat = new(pb.Wechat)
	}
	pbWechat.Id = wechat.Id
	pbWechat.Name = wechat.Name
	pbWechat.Appid = wechat.Appid
	pbWechat.Appsecret = wechat.Appsecret
	pbWechat.EncodingAesKey = wechat.EncodingAesKey
	pbWechat.Token = wechat.Token
	pbWechat.Flag = wechat.Flag
	pbWechat.Type = wechat.Type
	pbWechat.Pass = util.Int8ToBoolen(wechat.Pass)
	pbWechat.SaveInput = util.Int8ToBoolen(wechat.SaveInput)
	pbWechat.NeedSaveMen = util.Int8ToBoolen(wechat.NeedSaveMen)
	pbWechat.CreatedAt = wechat.CreatedAt.Unix()
	pbWechat.UpdatedAt = wechat.UpdatedAt.Unix()
	return pbWechat
}

func (this *WechatService) GetById(ctx context.Context, resq *pb.RequestById, resp *pb.Wechat) error {
	if resq.Id < 1 {
		return nil
	}
	wechat := dao.GetWechatServiceR().GetById(resq.Id)
	CopyWechatDaoToPb(wechat, resp)
	return nil
}

func (this *WechatService) GetAppId(ctx context.Context, resq *pb.RequestBySingleColumn, resp *pb.Wechat) error {
	if util.IsEmpty(resq.Column) {
		return nil
	}
	wechat := dao.GetWechatServiceR().GetByAppid(resq.Column)
	CopyWechatDaoToPb(wechat, resp)
	return nil
}

func (this *WechatService) GetByFlag(ctx context.Context, resq *pb.RequestBySingleColumn, resp *pb.Wechat) error {
	if util.IsEmpty(resq.Column) {
		return nil
	}
	wechat := dao.GetWechatServiceR().GetByFlag(resq.Column)
	CopyWechatDaoToPb(wechat, resp)
	return nil
}

func (this *WechatService) Insert(ctx context.Context, resq *pb.Wechat, resp *pb.Wechat) error {
	wechat := &dao.WechatModel{
		Name:           resq.Name,
		Appid:          resq.Appid,
		Appsecret:      resq.Appsecret,
		EncodingAesKey: resq.EncodingAesKey,
		Token:          resq.Token,
		Flag:           resq.Flag,
		Type:           resq.Type,
		Pass:           util.BoolenToInt8(resq.Pass),
		SaveInput:      util.BoolenToInt8(resq.SaveInput),
		NeedSaveMen:    util.BoolenToInt8(resq.NeedSaveMen),
		CreatedAt:      time.Unix(resq.CreatedAt, 0),
		UpdatedAt:      time.Unix(resq.UpdatedAt, 0),
	}
	_, err := dao.GetWechatServiceW().Insert(wechat)
	if err == nil {
		CopyWechatDaoToPb(wechat, resp)
	}
	return err
}

func (this *WechatService) Delete(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	ok := dao.GetWechatServiceW().DeleteById(resq.Id)
	if ok {
		resp.Success = true
		resp.Effect = 1
	}
	return nil
}
