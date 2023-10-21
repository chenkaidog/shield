package service

import (
	"context"
	"shield/account/internal/repos"
	"shield/account/model/domain"
	"shield/account/model/po"
	"shield/common/errs"
	"shield/common/utils/idgen"
)

func CreateUser(ctx context.Context, req *domain.UserCreateReq) (string, errs.Error) {
	userID := idgen.NewUUID()
	err := repos.CreateUser(ctx, &po.User{
		UserID:      userID,
		AccountID:   req.AccountID,
		Name:        req.Name,
		Gender:      string(req.Gender),
		Phone:       req.Phone,
		Email:       req.Email,
		Description: req.Description,
	})
	if err != nil {
		if errs.ErrorEqual(err, errs.DbDuplicateError) {
			return "", errs.UserDuplicateError
		}

		return "", err
	}

	return userID, nil
}

func UpdateUser(ctx context.Context, req *domain.UserUpdateReq) errs.Error {
	return repos.UpdateUser(ctx, &po.User{
		UserID:      req.UserID,
		Name:        req.Name,
		Gender:      string(req.Gender),
		Phone:       req.Phone,
		Email:       req.Email,
		Description: req.Description,
	})
}

func QueryUser(ctx context.Context, req *domain.UserQueryReq) (*domain.User, errs.Error) {
	var result *po.User
	var err errs.Error
	if req.UserID != "" {
		result, err = repos.SelectUserByID(ctx, req.UserID)
		if err != nil {
			return nil, err
		}
	}
	if req.AccountID != "" {
		result, err = repos.SelectUserByAccountID(ctx, req.AccountID)
		if err != nil {
			return nil, err
		}
	}

	if result == nil {
		return nil, nil
	}

	return &domain.User{
		UserID:      result.UserID,
		AccountID:   result.AccountID,
		Name:        result.Name,
		Gender:      domain.Gender(result.Gender),
		Phone:       result.Phone,
		Email:       result.Email,
		Description: result.Description,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}
