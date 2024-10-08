package logic

import (
	"context"

	"go-zero-learn/http/server-start/user/internal/svc"
	"go-zero-learn/http/server-start/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserReqResp, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info(req)
	resp = &types.UserReqResp{
		UserReq: *req,
		Status:  "OK",
	}
	return
}
