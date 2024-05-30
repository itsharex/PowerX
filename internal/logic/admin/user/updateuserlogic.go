package user

import (
	"PowerX/internal/model/origanzation"
	"PowerX/internal/types"
	"context"
	"github.com/pkg/errors"
	"time"

	"PowerX/internal/model"
	"PowerX/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *types.UpdateUserReply, err error) {
	user := origanzation.User{
		Model: model.Model{
			Id: req.Id,
		},
		Name:          req.Name,
		NickName:      req.NickName,
		Desc:          req.Desc,
		PositionID:    req.PositionId,
		JobTitle:      req.JobTitle,
		DepartmentId:  req.DepId,
		MobilePhone:   req.MobilePhone,
		Gender:        req.Gender,
		Email:         req.Email,
		ExternalEmail: req.ExternalEmail,
		Avatar:        req.Avatar,
		Password:      req.Password,
		Status:        req.Status,
	}

	if err = user.HashPassword(); err != nil {
		panic(errors.Wrap(err, "create user hash password failed"))
	}

	if err := l.svcCtx.PowerX.Organization.UpdateUserById(l.ctx, &user, req.Id); err != nil {
		return nil, err
	}

	// 根据职位更新角色
	if user.PositionID != 0 {
		codes, err := l.svcCtx.PowerX.Organization.FindUserPositionRoleCodes(l.ctx, user.Id)
		if err != nil {
			panic(err)
		}
		if _, err := l.svcCtx.PowerX.AdminAuthorization.Casbin.DeleteRolesForUser(user.Account); err != nil {
			panic(err)
		}
		if _, err := l.svcCtx.PowerX.AdminAuthorization.Casbin.AddRolesForUser(user.Account, codes); err != nil {
			panic(err)
		}
	}

	roles, _ := l.svcCtx.PowerX.AdminAuthorization.Casbin.GetRolesForUser(user.Account)

	return &types.UpdateUserReply{
		User: &types.User{
			Id:            user.Id,
			Account:       user.Account,
			Name:          user.Name,
			Email:         user.Email,
			MobilePhone:   user.MobilePhone,
			Gender:        user.Gender,
			NickName:      user.NickName,
			Desc:          user.Desc,
			Avatar:        user.Avatar,
			ExternalEmail: user.ExternalEmail,
			Roles:         roles,
			PositionId:    user.PositionID,
			JobTitle:      user.JobTitle,
			IsEnabled:     user.Status == origanzation.UserStatusEnabled,
			CreatedAt:     user.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}
