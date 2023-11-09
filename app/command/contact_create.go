package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/contact"
)

type ContactCreateCmd struct {
	Subject string `json:"subject" validate:"required,min=2,max=50"`
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required,min=10,max=500"`
}

type ContactCreateRes struct{}

type ContactCreateHandler cqrs.HandlerFunc[ContactCreateCmd, *ContactCreateRes]

func NewContactCreateHandler(factory contact.Factory, repo contact.Repository) ContactCreateHandler {
	return func(ctx context.Context, cmd ContactCreateCmd) (*ContactCreateRes, *i18np.Error) {
		return &ContactCreateRes{}, repo.Create(ctx, factory.New(contact.NewConfig{
			Subject: cmd.Subject,
			Email:   cmd.Email,
			Message: cmd.Message,
		}))
	}
}
