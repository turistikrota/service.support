package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportDeleteCmd struct {
	UUID     string `params:"uuid" validate:"required,object_id"`
	UserUUID string `json:"-"`
	UserName string `json:"-"`
}

type SupportDeleteRes struct{}

type SupportDeleteHandler cqrs.HandlerFunc[SupportDeleteCmd, *SupportDeleteRes]

func NewSupportDeleteHandler(repo support.Repo) SupportDeleteHandler {
	return func(ctx context.Context, cmd SupportDeleteCmd) (*SupportDeleteRes, *i18np.Error) {
		return &SupportDeleteRes{}, repo.Delete(ctx, cmd.UUID, support.WithUser{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
	}
}
