package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/support"
)

type SupportCreateCmd struct {
	UserUUID string `json:"-"`
	UserName string `json:"-"`
	Subject  string `json:"subject" validate:"required,min=3,max=100"`
	Message  string `json:"message" validate:"required,min=3,max=1000"`
}

type SupportCreateRes struct {
	UUID string `json:"uuid"`
}

type SupportCreateHandler cqrs.HandlerFunc[SupportCreateCmd, *SupportCreateRes]

func NewSupportCreateHandler(factory support.Factory, repo support.Repo) SupportCreateHandler {
	return func(ctx context.Context, cmd SupportCreateCmd) (*SupportCreateRes, *i18np.Error) {
		res, err := repo.Create(ctx, factory.New(support.NewConfig{
			User:    &support.User{UUID: cmd.UserUUID, Name: cmd.UserName},
			Subject: cmd.Subject,
			Message: cmd.Message,
		}))
		if err != nil {
			return nil, err
		}
		return &SupportCreateRes{
			UUID: res.UUID,
		}, nil
	}
}
