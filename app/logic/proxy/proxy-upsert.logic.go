package proxy

import (
	"context"

	"github.com/tradalab/rdms/app/dal/do"
	"github.com/tradalab/rdms/app/svc"
)

type ProxyUpsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProxyUpsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProxyUpsertLogic {
	return &ProxyUpsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProxyUpsertLogic) ProxyUpsertLogic(args *do.ProxyDO) (string, error) {
	err := l.svcCtx.GormMod.DB().WithContext(l.ctx).Save(args).Error
	if err != nil {
		return "", err
	}
	return args.Id.String(), nil
}
