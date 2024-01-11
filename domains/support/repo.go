package support

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WithUser struct {
	UUID string
	Name string
}

type Repo interface {
	// admin actions
	AdminFilter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
	AdminGet(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	AdminClose(ctx context.Context, uuid string) *i18np.Error
	AdminDelete(ctx context.Context, uuid string) *i18np.Error
	AdminAddMessage(ctx context.Context, uuid string, adminId string, message string) *i18np.Error
	AdminUpdate(ctx context.Context, uuid string, subject string, interests []string) *i18np.Error
	AdminRemoveMessage(ctx context.Context, uuid string, messageId string) *i18np.Error

	// user actions
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	AddMessage(ctx context.Context, uuid string, message string) *i18np.Error
	Close(ctx context.Context, uuid string, user WithUser) *i18np.Error
	Delete(ctx context.Context, uuid string, user WithUser) *i18np.Error
	Get(ctx context.Context, uuid string, user WithUser) (*Entity, *i18np.Error)
	Filter(ctx context.Context, user WithUser, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func NewRepo(collection *mongo.Collection, factory Factory) Repo {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}

func (r *repo) AdminFilter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	filters := r.filterToBson(filter)
	l, err := r.helper.GetListFilter(ctx, filters, r.listOptions(listConfig))
	if err != nil {
		return nil, err
	}
	filteredCount, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, bson.M{})
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filteredCount > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filteredCount,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) AdminGet(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	return nil, nil
}

func (r *repo) AdminClose(ctx context.Context, uuid string) *i18np.Error {
	return nil
}

func (r *repo) AdminDelete(ctx context.Context, uuid string) *i18np.Error {
	return nil
}

func (r *repo) AdminAddMessage(ctx context.Context, uuid string, adminId string, message string) *i18np.Error {
	return nil
}

func (r *repo) AdminUpdate(ctx context.Context, uuid string, subject string, interests []string) *i18np.Error {
	return nil
}

func (r *repo) AdminRemoveMessage(ctx context.Context, uuid string, messageId string) *i18np.Error {
	return nil
}

func (r *repo) Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error) {
	return nil, nil
}

func (r *repo) AddMessage(ctx context.Context, uuid string, message string) *i18np.Error {
	return nil
}

func (r *repo) Close(ctx context.Context, uuid string, user WithUser) *i18np.Error {
	return nil
}

func (r *repo) Delete(ctx context.Context, uuid string, user WithUser) *i18np.Error {
	return nil
}

func (r *repo) Get(ctx context.Context, uuid string, user WithUser) (*Entity, *i18np.Error) {
	return nil, nil
}

func (r *repo) Filter(ctx context.Context, user WithUser, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	return nil, nil
}

func (r *repo) listOptions(listConfig list.Config) *options.FindOptions {
	opts := options.Find()
	opts.SetSort(bson.M{
		fields.CreatedAt: -1,
	})
	opts.SetLimit(listConfig.Limit)
	opts.SetSkip(listConfig.Offset)
	return opts
}
