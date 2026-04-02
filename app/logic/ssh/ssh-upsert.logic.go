package ssh

import (
	"context"

	"github.com/tradalab/rdms/app/dal/do"
	"github.com/tradalab/rdms/app/svc"
)

type SshUpsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSshUpsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SshUpsertLogic {
	return &SshUpsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SshUpsertLogic) SshUpsertLogic(args *do.SshDO) (string, error) {
	err := l.svcCtx.GormMod.DB().WithContext(l.ctx).Save(args).Error
	if err != nil {
		return "", err
	}
	return args.Id.String(), nil
}
