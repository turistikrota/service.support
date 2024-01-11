package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportAdminRemoveMsgCmd struct {
	UUID      string `params:"uuid" validate:"required,object_id"`
	MessageId string `params:"message_id" validate:"required,uuid"`
}

type SupportAdminRemoveMsgRes struct{}

type SupportAdminRemoveMsgHandler cqrs.HandlerFunc[SupportAdminRemoveMsgCmd, *SupportAdminRemoveMsgRes]

func NewSupportAdminRemoveMsgHandler(repo support.Repo) SupportAdminRemoveMsgHandler {
	return func(ctx context.Context, cmd SupportAdminRemoveMsgCmd) (*SupportAdminRemoveMsgRes, *i18np.Error) {
		return &SupportAdminRemoveMsgRes{}, repo.AdminRemoveMessage(ctx, cmd.UUID, cmd.MessageId)
	}
}
