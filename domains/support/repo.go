package support

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/google/uuid"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	AdminGet(ctx context.Context, uuid string) (*Entity, bool, *i18np.Error)
	AdminClose(ctx context.Context, uuid string) *i18np.Error
	AdminDelete(ctx context.Context, uuid string) *i18np.Error
	AdminAddMessage(ctx context.Context, uuid string, adminId string, message string) *i18np.Error
	AdminUpdate(ctx context.Context, uuid string, subject string) *i18np.Error
	AdminRemoveMessage(ctx context.Context, uuid string, messageId string) *i18np.Error

	// user actions
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	AddMessage(ctx context.Context, uuid string, message string, user WithUser) *i18np.Error
	Close(ctx context.Context, uuid string, user WithUser) *i18np.Error
	Delete(ctx context.Context, uuid string, user WithUser) *i18np.Error
	Get(ctx context.Context, uuid string, user WithUser) (*Entity, bool, *i18np.Error)
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

func (r *repo) AdminGet(ctx context.Context, uuid string) (*Entity, bool, *i18np.Error) {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return nil, false, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	res, notFound, _err := r.helper.GetFilter(ctx, filter)
	if _err != nil {
		return nil, false, _err
	}
	if notFound {
		return nil, true, nil
	}
	return *res, false, nil
}

func (r *repo) AdminClose(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsUserClosed: false,
			fields.State:        States.Closed,
			fields.ClosedAt:     time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) AdminDelete(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsUserClosed: false,
			fields.State:        States.Deleted,
			fields.ClosedAt:     nil,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) AdminAddMessage(ctx context.Context, supportId string, adminId string, message string) *i18np.Error {
	id, err := mongo2.TransformId(supportId)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$addToSet": bson.M{
			fields.Messages: bson.M{
				messageFields.UUID:         uuid.New(),
				messageFields.InterestUUID: adminId,
				messageFields.Text:         message,
				messageFields.IsAdmin:      true,
				messageFields.IsDeleted:    false,
				messageFields.Date:         time.Now(),
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) AdminUpdate(ctx context.Context, uuid string, subject string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.Subject: subject,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) AdminRemoveMessage(ctx context.Context, supportId string, messageId string) *i18np.Error {
	id, err := mongo2.TransformId(supportId)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                      id,
		messageField(messageFields.UUID): messageId,
	}
	update := bson.M{
		"$pull": bson.M{
			fields.Messages: bson.M{
				messageFields.UUID: messageId,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error) {
	res, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	entity.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}

func (r *repo) AddMessage(ctx context.Context, supportId string, message string, user WithUser) *i18np.Error {
	id, err := mongo2.TransformId(supportId)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                id,
		userField(userFields.UUID): user.UUID,
		userField(userFields.Name): user.Name,
	}
	update := bson.M{
		"$addToSet": bson.M{
			fields.Messages: bson.M{
				messageFields.UUID:         uuid.New(),
				messageFields.InterestUUID: user.UUID,
				messageFields.Text:         message,
				messageFields.IsAdmin:      false,
				messageFields.IsDeleted:    false,
				messageFields.Date:         time.Now(),
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Close(ctx context.Context, uuid string, user WithUser) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                id,
		userField(userFields.UUID): user.UUID,
		userField(userFields.Name): user.Name,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsUserClosed: true,
			fields.State:        States.Closed,
			fields.ClosedAt:     time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Delete(ctx context.Context, uuid string, user WithUser) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                id,
		userField(userFields.UUID): user.UUID,
		userField(userFields.Name): user.Name,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsUserClosed: true,
			fields.State:        States.Deleted,
			fields.ClosedAt:     nil,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Get(ctx context.Context, uuid string, user WithUser) (*Entity, bool, *i18np.Error) {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return nil, false, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                id,
		userField(userFields.UUID): user.UUID,
		userField(userFields.Name): user.Name,
	}
	res, notFound, _err := r.helper.GetFilter(ctx, filter)
	if _err != nil {
		return nil, false, _err
	}
	if notFound {
		return nil, true, nil
	}
	return *res, false, nil
}

func (r *repo) Filter(ctx context.Context, user WithUser, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	filters := r.filterToBson(filter)
	l, err := r.helper.GetListFilter(ctx, filters, r.listOptions(listConfig))
	if err != nil {
		return nil, err
	}
	filteredCount, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, bson.M{
		userField(userFields.UUID): user.UUID,
		userField(userFields.Name): user.Name,
	})
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

func (r *repo) listOptions(listConfig list.Config) *options.FindOptions {
	opts := options.Find()
	opts.SetSort(bson.M{
		fields.CreatedAt: -1,
	})
	opts.SetLimit(listConfig.Limit)
	opts.SetSkip(listConfig.Offset)
	return opts
}
