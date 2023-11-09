package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/contact"
)

type ContactReadCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type ContactReadRes struct{}

type ContactReadHandler cqrs.HandlerFunc[ContactReadCmd, *ContactReadRes]

func NewContactReadHandler(factory contact.Factory, repo contact.Repository) ContactReadHandler {
	return func(ctx context.Context, cmd ContactReadCmd) (*ContactReadRes, *i18np.Error) {
		return &ContactReadRes{}, repo.Read(ctx, cmd.UUID)
	}
}
