package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportAdminAddMsgCmd struct {
	UUID     string `json:"-"`
	UserUUID string `json:"-"`
	Message  string `json:"message" validate:"required,min=10,max=500"`
}

type SupportAdminAddMsgRes struct{}

type SupportAdminAddMsgHandler cqrs.HandlerFunc[SupportAdminAddMsgCmd, *SupportAdminAddMsgRes]

func NewSupportAdminAddMsgHandler(repo support.Repo) SupportAdminAddMsgHandler {
	return func(ctx context.Context, cmd SupportAdminAddMsgCmd) (*SupportAdminAddMsgRes, *i18np.Error) {
		return &SupportAdminAddMsgRes{}, repo.AdminAddMessage(ctx, cmd.UUID, cmd.UserUUID, cmd.Message)
	}
}
