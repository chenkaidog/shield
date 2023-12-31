package handler

import (
	"context"
	"shield/account/internal/handler/service"
	"shield/account/internal/model/domain"
	"shield/account/kitex_gen/kaidog/shield/account"
	"shield/common/errs"

	"github.com/apache/thrift/lib/go/thrift"
)

func Login(ctx context.Context, req *account.AccountLoginReq) (*account.AccountLoginResp, errs.Error) {
	accountID, err := service.Login(ctx, &domain.LoginReq{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		IPv4:     req.GetIpv4(),
		Device:   req.GetDevice(),
	})
	if err != nil {
		return nil, err
	}

	resp := account.NewAccountLoginResp()
	resp.SetAccountID(&accountID)

	return resp, nil
}

func QueryLoginRecord(ctx context.Context, req *account.LoginRecordQueryReq) (*account.LoginRecordQueryResp, errs.Error) {
	resultList, total, err := service.QueryLoginRecord(ctx, &domain.LoginRecordQueryReq{
		AccountID: req.GetAccountID(),
		Page:      req.GetPage(),
		Size:      req.GetSize(),
	})
	if err != nil {
		return nil, err
	}

	resp := account.NewLoginRecordQueryResp()
	resp.SetPage(thrift.Int64Ptr(req.GetPage()))
	resp.SetSize(thrift.Int64Ptr(req.GetSize()))
	resp.SetTotal(thrift.Int64Ptr(total))

	for _, record := range resultList {
		var status account.LoginStatus
		switch record.Status {
		case domain.LoginStatusSuccess:
			status = account.LoginStatus_success
		case domain.LoginStatusFail:
			status = account.LoginStatus_fail
		}
		resp.RecordList = append(resp.RecordList, &account.LoginRecord{
			AccountID: record.AccountID,
			LoginAt:   record.LoginAt.Unix(),
			Ipv4:      record.IPv4,
			Device:    record.Device,
			Reason:    record.Reason,
			Status:    status,
		})
	}

	return resp, nil
}
