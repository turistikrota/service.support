package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportAdminGetQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type SupportAdminGetRes struct {
	Detail support.AdminDto
}

type SupportAdminGetHandler cqrs.HandlerFunc[SupportAdminGetQuery, *SupportAdminGetRes]

func NewSupportAdminGetHandler(repo support.Repo) SupportAdminGetHandler {
	return func(ctx context.Context, query SupportAdminGetQuery) (*SupportAdminGetRes, *i18np.Error) {
		res, err := repo.AdminGet(ctx, query.UUID)
		if err != nil {
			return nil, err
		}
		return &SupportAdminGetRes{
			Detail: res.ToAdmin(),
		}, nil
	}
}
