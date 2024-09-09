// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: user.proto

package server

import (
	"context"

	"go-zero-learn/rpc/server/internal/logic/user"
	"go-zero-learn/rpc/server/internal/svc"
	"go-zero-learn/rpc/server/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Create(ctx context.Context, in *user.UserReq) (*user.UserResp, error) {
	l := userlogic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}
