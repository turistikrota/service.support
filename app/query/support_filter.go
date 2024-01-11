package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.support/domains/support"
	"github.com/turistikrota/service.support/pkg/utils"
)

type SupportFilterQuery struct {
	*utils.Pagination
	*support.FilterEntity
	UserUUID string `json:"-"`
	UserName string `json:"-"`
}

type SupportFilterRes struct {
	List *list.Result[support.ListDto]
}

type SupportFilterHandler cqrs.HandlerFunc[SupportFilterQuery, *SupportFilterRes]

func NewSupportFilterHandler(repo support.Repo) SupportFilterHandler {
	return func(ctx context.Context, query SupportFilterQuery) (*SupportFilterRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.Filter(ctx, support.WithUser{
			UUID: query.UserUUID,
			Name: query.UserName,
		}, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]support.ListDto, len(res.List))
		for i, e := range res.List {
			li[i] = e.ToList()
		}
		return &SupportFilterRes{
			List: &list.Result[support.ListDto]{
				List:          li,
				Total:         res.Total,
				FilteredTotal: res.FilteredTotal,
				Page:          res.Page,
				IsNext:        res.IsNext,
				IsPrev:        res.IsPrev,
			},
		}, nil
	}
}
