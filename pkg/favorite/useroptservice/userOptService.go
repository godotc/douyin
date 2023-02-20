// Code generated by goctl. DO NOT EDIT.
// Source: UserOptService.proto

package useroptservice

import (
	"context"
	userOptPb2 "douyin/pkg/favorite/userOptPb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Favorite                 = userOptPb2.Favorite
	GetUserFavoriteReq       = userOptPb2.GetUserFavoriteReq
	GetUserFavoriteResp      = userOptPb2.GetUserFavoriteResp
	UpdateFavoriteStatusReq  = userOptPb2.UpdateFavoriteStatusReq
	UpdateFavoriteStatusResp = userOptPb2.UpdateFavoriteStatusResp

	UserOptService interface {
		// -----------------------userFavoriteList-----------------------
		GetUserFavorite(ctx context.Context, in *GetUserFavoriteReq, opts ...grpc.CallOption) (*GetUserFavoriteResp, error)
		UpdateFavoriteStatus(ctx context.Context, in *UpdateFavoriteStatusReq, opts ...grpc.CallOption) (*UpdateFavoriteStatusResp, error)
	}

	defaultUserOptService struct {
		cli zrpc.Client
	}
)

func NewUserOptService(cli zrpc.Client) UserOptService {
	return &defaultUserOptService{
		cli: cli,
	}
}

// -----------------------userFavoriteList-----------------------
func (m *defaultUserOptService) GetUserFavorite(ctx context.Context, in *GetUserFavoriteReq, opts ...grpc.CallOption) (*GetUserFavoriteResp, error) {
	client := userOptPb2.NewUserOptServiceClient(m.cli.Conn())
	return client.GetUserFavorite(ctx, in, opts...)
}

func (m *defaultUserOptService) UpdateFavoriteStatus(ctx context.Context, in *UpdateFavoriteStatusReq, opts ...grpc.CallOption) (*UpdateFavoriteStatusResp, error) {
	client := userOptPb2.NewUserOptServiceClient(m.cli.Conn())
	return client.UpdateFavoriteStatus(ctx, in, opts...)
}
