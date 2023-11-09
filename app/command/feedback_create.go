package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.support/domains/feedback"
)

type FeedbackCreateCmd struct {
	OS      string         `json:"os" validate:"required,min=2,max=50"`
	Version string         `json:"version" validate:"required,min=2,max=50"`
	User    *feedback.User `json:"-"`
	Message string         `json:"message" validate:"required,min=10,max=500"`
}

type FeedbackCreateRes struct{}

type FeedbackCreateHandler cqrs.HandlerFunc[FeedbackCreateCmd, *FeedbackCreateRes]

func NewFeedbackCreateHandler(factory feedback.Factory, repo feedback.Repository) FeedbackCreateHandler {
	return func(ctx context.Context, cmd FeedbackCreateCmd) (*FeedbackCreateRes, *i18np.Error) {
		return &FeedbackCreateRes{}, repo.Create(ctx, factory.New(feedback.NewConfig{
			OS:      cmd.OS,
			Version: cmd.Version,
			User:    cmd.User,
			Message: cmd.Message,
		}))
	}
}
