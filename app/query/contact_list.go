package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.support/domains/contact"
	"github.com/turistikrota/service.support/pkg/utils"
)

type ContactListQuery struct {
	*utils.Pagination
	*contact.FilterEntity
}

type ContactListRes struct {
	*list.Result[*contact.Entity]
}

type ContactListHandler cqrs.HandlerFunc[ContactListQuery, *ContactListRes]

func NewContactListHandler(repo contact.Repository) ContactListHandler {
	return func(ctx context.Context, query ContactListQuery) (*ContactListRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		list, err := repo.List(ctx, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		return &ContactListRes{
			Result: list,
		}, nil
	}
}
