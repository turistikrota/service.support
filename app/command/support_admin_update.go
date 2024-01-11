package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportAdminUpdateCmd struct {
	UUID    string `json:"-"`
	Subject string `json:"subject" validate:"required,min=10,max=100"`
}

type SupportAdminUpdateRes struct{}

type SupportAdminUpdateHandler cqrs.HandlerFunc[SupportAdminUpdateCmd, *SupportAdminUpdateRes]

func NewSupportAdminUpdateHandler(repo support.Repo) SupportAdminUpdateHandler {
	return func(ctx context.Context, cmd SupportAdminUpdateCmd) (*SupportAdminUpdateRes, *i18np.Error) {
		return &SupportAdminUpdateRes{}, repo.AdminUpdate(ctx, cmd.UUID, cmd.Subject)
	}
}
