package tls

import (
	"context"

	"github.com/tradalab/rdms/app/dal/do"
	"github.com/tradalab/rdms/app/svc"
)

type TlsUpsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTlsUpsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TlsUpsertLogic {
	return &TlsUpsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TlsUpsertLogic) TlsUpsertLogic(args *do.TlsDO) (string, error) {
	err := l.svcCtx.GormMod.DB().WithContext(l.ctx).Save(args).Error
	if err != nil {
		return "", err
	}
	return args.Id.String(), nil
}
