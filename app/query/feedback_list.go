package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.support/domains/feedback"
	"github.com/turistikrota/service.support/pkg/utils"
)

type FeedbackListQuery struct {
	*utils.Pagination
	*feedback.FilterEntity
}

type FeedbackListRes struct {
	*list.Result[*feedback.Entity]
}

type FeedbackListHandler cqrs.HandlerFunc[FeedbackListQuery, *FeedbackListRes]

func NewFeedbackListHandler(repo feedback.Repository) FeedbackListHandler {
	return func(ctx context.Context, query FeedbackListQuery) (*FeedbackListRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		list, err := repo.List(ctx, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		return &FeedbackListRes{
			Result: list,
		}, nil
	}
}
