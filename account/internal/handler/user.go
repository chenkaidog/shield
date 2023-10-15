package handler

import (
	"context"
	"shield/account/internal/handler/service"
	"shield/account/kitex_gen/kaidog/shield/account"
	"shield/account/model/domain"
	"shield/common/errs"

	"github.com/apache/thrift/lib/go/thrift"
)

func CreateUser(ctx context.Context, req *account.UserCreateReq) (resp *account.UserCreateResp, err errs.Error) {
	var gender domain.Gender
	switch req.GetGender() {
	case account.Gender_male:
		gender = domain.GenderMale
	case account.Gender_female:
		gender = domain.GenderFemale
	case account.Gender_others:
		gender = domain.GenderOthers
	}

	userID, err := service.CreateUser(ctx, &domain.UserCreateReq{
		AccountID:   req.GetAccountID(),
		Name:        req.GetName(),
		Gender:      gender,
		Phone:       req.GetPhone(),
		Email:       req.GetEmail(),
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}

	resp = account.NewUserCreateResp()
	resp.SetUserID(thrift.StringPtr(userID))

	return resp, nil
}

func UpdateUser(ctx context.Context, req *account.UserUpdateReq) (resp *account.UserUpdateResp, err errs.Error) {
	var gender domain.Gender
	switch req.GetGender() {
	case account.Gender_male:
		gender = domain.GenderMale
	case account.Gender_female:
		gender = domain.GenderFemale
	case account.Gender_others:
		gender = domain.GenderOthers
	}

	err = service.UpdateUser(ctx, &domain.UserUpdateReq{
		UserID:      req.GetUserID(),
		Name:        req.GetName(),
		Gender:      gender,
		Phone:       req.GetPhone(),
		Email:       req.GetEmail(),
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}

	resp = account.NewUserUpdateResp()
	return resp, nil
}

func QueryUser(ctx context.Context, req *account.UserQueryReq) (resp *account.UserQueryResp, err errs.Error) {
	result, err := service.QueryUser(ctx, &domain.UserQueryReq{
		UserID:    req.GetUserID(),
		AccountID: req.GetAccountID(),
	})
	if err != nil {
		return nil, err
	}

	resp = account.NewUserQueryResp()
	if result != nil {
		var gender account.Gender
		switch result.Gender {
		case domain.GenderMale:
			gender = account.Gender_male
		case domain.GenderFemale:
			gender = account.Gender_female
		case domain.GenderOthers:
			gender = account.Gender_others
		}

		resp.SetUser(&account.User{
			AccountID:   result.AccountID,
			UserID:      result.UserID,
			Name:        result.Name,
			Gender:      gender,
			Phone:       result.Phone,
			Email:       result.Email,
			Description: result.Description,
			CreatedAt:   result.CreatedAt.Unix(),
			UpdatedAt:   result.UpdatedAt.Unix(),
		})
	}

	return resp, nil
}
