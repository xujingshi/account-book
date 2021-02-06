package service

import (
	"account-book/internal/model"
	"context"
)

func (s *Service) AddChargeDetails(ctx context.Context, mid int64, req *model.AddChargeDetailsReq) (rsp *model.NoReply, err error) {
	rsp = new(model.NoReply)
	err = s.dao.AddChargeDetails(ctx, mid, req.Details, req.AutoImport, req.Status)
	return
}

func (s *Service) GetDetailList(ctx context.Context, mid int64, req *model.GetDetailListReq) (rsp *model.GetDetailListRsp, err error) {
	rsp = new(model.GetDetailListRsp)
	details, err := s.dao.GetChargeDetails(ctx, mid, req.Year, req.Month, req.Day, req.Module, req.Status)
	if err != nil {
		return
	}
	rsp.Details = details
	return
}