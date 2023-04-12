package logic

import (
	"context"

	"myapi/internal/svc"
	"myapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyapiLogic {
	return &MyapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyapiLogic) Myapi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
