package logic

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"os"

	"go-zero-learn/http_download/internal/svc"
	"go-zero-learn/http_download/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownLoadReq, w http.ResponseWriter) error {
	body, err := os.ReadFile(req.Filename)
	if err != nil {
		httpx.Error(w, err)
		return err
	}

	_, err = w.Write(body)
	if err != nil {
		return err
	}

	return nil
}
