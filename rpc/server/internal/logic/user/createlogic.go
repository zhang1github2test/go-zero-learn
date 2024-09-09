package userlogic

import (
	"context"

	"go-zero-learn/rpc/server/internal/svc"
	"go-zero-learn/rpc/server/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.UserReq) (*user.UserResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserResp{}, nil
}
