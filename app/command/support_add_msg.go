package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportAddMsgCmd struct {
	UUID     string `json:"-"`
	UserUUID string `json:"-"`
	UserName string `json:"-"`
	Message  string `json:"message" validate:"required,min=10,max=500"`
}

type SupportAddMsgRes struct{}

type SupportAddMsgHandler cqrs.HandlerFunc[SupportAddMsgCmd, *SupportAddMsgRes]

func NewSupportAddMsgHandler(repo support.Repo) SupportAddMsgHandler {
	return func(ctx context.Context, cmd SupportAddMsgCmd) (*SupportAddMsgRes, *i18np.Error) {
		return &SupportAddMsgRes{}, repo.AddMessage(ctx, cmd.UUID, cmd.Message, support.WithUser{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
	}
}
