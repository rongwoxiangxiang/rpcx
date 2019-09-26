package service

import (
	"context"
	"rpc/dao"
	"rpc/pb"
	"rpc/util"
	"time"
	"unsafe"
)

type RecordService struct {
}

func CopyRecordDaoToPb(record *dao.RecordModel, pbRecord *pb.Record) *pb.Record {
	if pbRecord == nil {
		pbRecord = new(pb.Record)
	}
	pbRecord.Id = record.Id
	pbRecord.Wid = record.Wid
	pbRecord.Wuid = record.Wuid
	pbRecord.Type = record.Type
	pbRecord.Content = record.Content
	pbRecord.CreatedAt = record.CreatedAt.Unix()
	return pbRecord
}

func (this *RecordService) GetById(ctx context.Context, resq *pb.RequestById, resp *pb.Record) error {
	if resq.Id < 1 {
		return nil
	}
	record := dao.GetRecordServiceR().GetById(resq.Id)
	CopyRecordDaoToPb(record, resp)
	return nil
}

func (this *RecordService) LimitByWid(ctx context.Context, resq *pb.RequestList, resp *pb.RecordList) error {
	wid, ok := resq.Params["wid"]
	if !ok {
		return nil
	}
	widInt64 := util.StringToInt64(wid)
	records := dao.GetRecordServiceR().LimitByWid(
		widInt64,
		*(*int)(unsafe.Pointer(&resq.Index)),
		*(*int)(unsafe.Pointer(&resq.Limit)))
	for _, record := range records {
		resp.Records = append(resp.Records, CopyRecordDaoToPb(record, nil))
	}
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	resp.Total = dao.GetRecordServiceR().Count(&dao.RecordModel{Wid: widInt64})
	return nil
}

func (this *RecordService) LimitByWuid(ctx context.Context, resq *pb.RequestList, resp *pb.RecordList) error {
	wuid, ok := resq.Params["wuid"]
	if !ok {
		return nil
	}
	wuidInt64 := util.StringToInt64(wuid)
	records := dao.GetRecordServiceR().LimitByWuid(
		wuidInt64,
		*(*int)(unsafe.Pointer(&resq.Index)),
		*(*int)(unsafe.Pointer(&resq.Limit)))
	for _, record := range records {
		resp.Records = append(resp.Records, CopyRecordDaoToPb(record, nil))
	}
	resp.Limit = resq.Limit
	resp.Index = resq.Index
	resp.Total = dao.GetRecordServiceR().Count(&dao.RecordModel{Wuid: wuidInt64})
	return nil
}

func (this *RecordService) Insert(ctx context.Context, resq *pb.Record, resp *pb.Record) error {
	record := &dao.RecordModel{
		Wid:       resq.Wid,
		Wuid:      resq.Wuid,
		Type:      resq.Type,
		Content:   resq.Content,
		CreatedAt: time.Unix(resq.CreatedAt, 0),
	}
	_, err := dao.GetRecordServiceW().Insert(record)
	if err == nil {
		CopyRecordDaoToPb(record, resp)
	}
	return err
}
