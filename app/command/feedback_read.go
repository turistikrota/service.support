package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/feedback"
)

type FeedbackReadCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type FeedbackReadRes struct{}

type FeedbackReadHandler cqrs.HandlerFunc[FeedbackReadCmd, *FeedbackReadRes]

func NewFeedbackReadHandler(factory feedback.Factory, repo feedback.Repository) FeedbackReadHandler {
	return func(ctx context.Context, cmd FeedbackReadCmd) (*FeedbackReadRes, *i18np.Error) {
		return &FeedbackReadRes{}, repo.Read(ctx, cmd.UUID)
	}
}
