package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportCloseCmd struct {
	UUID     string `params:"uuid" validate:"required,object_id"`
	UserUUID string `json:"-"`
	UserName string `json:"-"`
}

type SupportCloseRes struct{}

type SupportCloseHandler cqrs.HandlerFunc[SupportCloseCmd, *SupportCloseRes]

func NewSupportCloseHandler(repo support.Repo) SupportCloseHandler {
	return func(ctx context.Context, cmd SupportCloseCmd) (*SupportCloseRes, *i18np.Error) {
		return &SupportCloseRes{}, repo.Close(ctx, cmd.UUID, support.WithUser{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
	}
}
