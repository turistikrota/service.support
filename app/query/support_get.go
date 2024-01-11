package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportGetQuery struct {
	UUID     string `params:"uuid" validate:"required,object_id"`
	UserUUID string `json:"-"`
	UserName string `json:"-"`
}

type SupportGetRes struct {
	Detail support.DetailDto
}

type SupportGetHandler cqrs.HandlerFunc[SupportGetQuery, *SupportGetRes]

func NewSupportGetHandler(repo support.Repo) SupportGetHandler {
	return func(ctx context.Context, query SupportGetQuery) (*SupportGetRes, *i18np.Error) {
		res, err := repo.Get(ctx, query.UUID, support.WithUser{
			UUID: query.UserUUID,
			Name: query.UserName,
		})
		if err != nil {
			return nil, err
		}
		return &SupportGetRes{
			Detail: res.ToDetail(),
		}, nil
	}
}
