package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.support/domains/support"
	"github.com/turistikrota/service.support/pkg/utils"
)

type SupportAdminFilterQuery struct {
	*utils.Pagination
	*support.FilterEntity
}

type SupportAdminFilterRes struct {
	List *list.Result[support.AdminListDto]
}

type SupportAdminFilterHandler cqrs.HandlerFunc[SupportAdminFilterQuery, *SupportAdminFilterRes]

func NewSupportAdminFilterHandler(repo support.Repo) SupportAdminFilterHandler {
	return func(ctx context.Context, query SupportAdminFilterQuery) (*SupportAdminFilterRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.AdminFilter(ctx, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]support.AdminListDto, len(res.List))
		for i, e := range res.List {
			li[i] = e.ToAdminList()
		}
		return &SupportAdminFilterRes{
			List: &list.Result[support.AdminListDto]{
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
