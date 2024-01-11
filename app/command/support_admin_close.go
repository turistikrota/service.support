package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportAdminCloseCmd struct {
	UUID string `json:"-"`
}

type SupportAdminCloseRes struct{}

type SupportAdminCloseHandler cqrs.HandlerFunc[SupportAdminCloseCmd, *SupportAdminCloseRes]

func NewSupportAdminCloseHandler(repo support.Repo) SupportAdminCloseHandler {
	return func(ctx context.Context, cmd SupportAdminCloseCmd) (*SupportAdminCloseRes, *i18np.Error) {
		return &SupportAdminCloseRes{}, repo.AdminClose(ctx, cmd.UUID)
	}
}
