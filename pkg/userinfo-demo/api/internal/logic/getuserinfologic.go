package logic

import (
	"context"

	"douyin/pkg/userinfo-demo/api/internal/svc"
	"douyin/pkg/userinfo-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoRes, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserInfoRes{
		StatusCode: 0,
		StatusMsg:  "",
		User: types.User{
			Id:            "234",
			Name:          "TBH",
			FollowCount:   11,
			FollowerCount: 12,
			IsFollow:      false,
		},
	}
	return
}