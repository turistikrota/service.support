package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportAdminDeleteCmd struct {
	UUID string `json:"-"`
}

type SupportAdminDeleteRes struct{}

type SupportAdminDeleteHandler cqrs.HandlerFunc[SupportAdminDeleteCmd, *SupportAdminDeleteRes]

func NewSupportAdminDeleteHandler(repo support.Repo) SupportAdminDeleteHandler {
	return func(ctx context.Context, cmd SupportAdminDeleteCmd) (*SupportAdminDeleteRes, *i18np.Error) {
		return &SupportAdminDeleteRes{}, repo.AdminDelete(ctx, cmd.UUID)
	}
}
