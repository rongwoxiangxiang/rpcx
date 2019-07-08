package service

import (
	"context"
	"github.com/jinzhu/copier"
	"rpc/lib/models"
	"time"
)

type CheckinServiceImpl struct{}

func (h CheckinServiceImpl) GetById(ctx context.Context, checkin *Checkin) (*Checkin, error) {
	return nil, nil
}

func (h CheckinServiceImpl) List(ctx context.Context, in *CheckinQuery) (*Checkinlist, error) {
	return nil, nil
}

func (h CheckinServiceImpl) Insert(ctx context.Context, checkin *Checkin) (*Checkin, error) {
	checkinModel := new(models.CheckinModel)
	copier.Copy(checkinModel, checkin)
	if checkin.Lastcheckin != 0 {
		checkinModel.Lastcheckin = time.Unix(checkin.Lastcheckin, 0)
	}
	insertId, err := checkinModel.Insert(checkinModel)
	checkin.Id = insertId
	return checkin, err
}
