package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-learn/http/server-start/user/internal/svc"
	"go-zero-learn/http/server-start/user/internal/types"
	"time"
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
	start := time.Now()
	time.Sleep(time.Second * 2)
	fmt.Printf("耗时:%v\n", time.Since(start))
	// todo: add your logic here and delete this line
	resp = &types.UserReqResp{
		UserReq: *req,
		Status:  "ok",
	}
	return
}
