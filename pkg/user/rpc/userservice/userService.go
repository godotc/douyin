// Code generated by goctl. DO NOT EDIT.
// Source: UserService.proto

package userservice

import (
	"context"

	"douyin/pkg/user/rpc/userInfoPb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthsInfoReq  = userInfoPb.AuthsInfoReq
	AuthsInfoResp = userInfoPb.AuthsInfoResp
	LoginReq      = userInfoPb.LoginReq
	LoginResp     = userInfoPb.LoginResp
	RegisterReq   = userInfoPb.RegisterReq
	RegisterResp  = userInfoPb.RegisterResp
	User          = userInfoPb.User
	UserInfoReq   = userInfoPb.UserInfoReq
	UserInfoResp  = userInfoPb.UserInfoResp

	UserService interface {
		// -----------------------user-----------------------
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Info(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

// -----------------------user-----------------------
func (m *defaultUserService) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := userInfoPb.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := userInfoPb.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserService) Info(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := userInfoPb.NewUserServiceClient(m.cli.Conn())
	return client.Info(ctx, in, opts...)
}