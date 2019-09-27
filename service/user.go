package service

import (
	"context"
	"rpc/dao"
	"rpc/pb"
	"rpc/util"
	"time"
	"unsafe"
)

type WechatUserService struct {
}

func CopyWechatUserDaoToPb(record *dao.WechatUserModel, pbWechatUser *pb.WechatUser) *pb.WechatUser {
	if pbWechatUser == nil {
		pbWechatUser = new(pb.WechatUser)
	}
	pbWechatUser.Id = record.Id
	pbWechatUser.Wid = record.Wid
	pbWechatUser.UserId = record.UserId
	pbWechatUser.Openid = record.Openid
	pbWechatUser.Nickname = record.Nickname
	pbWechatUser.Sex = int32(record.Sex)
	pbWechatUser.Province = record.Province
	pbWechatUser.City = record.City
	pbWechatUser.Country = record.Country
	pbWechatUser.Language = record.Language
	pbWechatUser.Headimgurl = record.Headimgurl
	pbWechatUser.CreatedAt = record.CreatedAt.Unix()
	pbWechatUser.UpdatedAt = record.UpdatedAt.Unix()
	return pbWechatUser
}

func (this *WechatUserService) GetById(ctx context.Context, resq *pb.RequestById, resp *pb.WechatUser) error {
	if resq.Id < 1 {
		return nil
	}
	wechatUser := dao.GetWechatUserServiceR().GetById(resq.Id)
	CopyWechatUserDaoToPb(wechatUser, resp)
	return nil
}

func (this *WechatUserService) LimitByWid(ctx context.Context, resq *pb.RequestList, resp *pb.WechatUserList) error {
	wid, ok := resq.Params["wid"]
	if !ok {
		return nil
	}
	widInt64 := util.StringToInt64(wid)
	users := dao.GetWechatUserServiceR().LimitByWid(
		widInt64,
		*(*int)(unsafe.Pointer(&resq.Index)),
		*(*int)(unsafe.Pointer(&resq.Limit)))
	for _, user := range users {
		resp.WechatUsers = append(resp.WechatUsers, CopyWechatUserDaoToPb(user, nil))
	}
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	resp.Total = dao.GetWechatUserServiceR().Count(&dao.WechatUserModel{Wid: widInt64})
	return nil
}

func (this *WechatUserService) LimitByWidAndOpenid(ctx context.Context, resq *pb.RequestOneByWidOpenid, resp *pb.WechatUser) error {
	users := dao.GetWechatUserServiceR().GetByWidAndOpenid(resq.Wid, resq.Openid)
	CopyWechatUserDaoToPb(users, resp)
	return nil
}

func (this *WechatUserService) Insert(ctx context.Context, resq *pb.WechatUser, resp *pb.WechatUser) error {
	record := &dao.WechatUserModel{
		Wid:        resq.Wid,
		UserId:     resq.UserId,
		Openid:     resq.Openid,
		Nickname:   resq.Nickname,
		Sex:        *(*int8)(unsafe.Pointer(&resq.Sex)),
		Province:   resq.Province,
		City:       resq.City,
		Country:    resq.Country,
		Language:   resq.Language,
		Headimgurl: resq.Headimgurl,
		CreatedAt:  time.Unix(resq.CreatedAt, 0),
		UpdatedAt:  time.Unix(resq.UpdatedAt, 0),
	}
	_, err := dao.GetWechatUserServiceW().Insert(record)
	if err == nil {
		CopyWechatUserDaoToPb(record, resp)
	}
	return err
}

func (this *WechatUserService) Update(ctx context.Context, resq *pb.WechatUser, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	rows, err := dao.GetWechatUserServiceW().Update(&dao.WechatUserModel{
		Id:         resq.Id,
		Wid:        resq.Wid,
		UserId:     resq.UserId,
		Openid:     resq.Openid,
		Nickname:   resq.Nickname,
		Sex:        *(*int8)(unsafe.Pointer(&resq.Sex)),
		Province:   resq.Province,
		City:       resq.City,
		Country:    resq.Country,
		Language:   resq.Language,
		Headimgurl: resq.Headimgurl,
		CreatedAt:  time.Unix(resq.CreatedAt, 0),
		UpdatedAt:  time.Unix(resq.UpdatedAt, 0),
	})
	if err == nil {
		resp.Success = true
		resp.Effect = rows
	}
	return err
}

func (this *WechatUserService) Delete(ctx context.Context, resq *pb.RequestById, resp *pb.ResponseEffect) error {
	if resq.Id < 0 {
		return nil
	}
	ok := dao.GetWechatUserServiceW().DeleteById(resq.Id)
	if ok {
		resp.Success = true
		resp.Effect = 1
	}
	return nil
}
