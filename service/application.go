package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.support/app"
	"github.com/turistikrota/service.support/app/command"
	"github.com/turistikrota/service.support/app/query"
	"github.com/turistikrota/service.support/config"
	"github.com/turistikrota/service.support/domains/contact"
	"github.com/turistikrota/service.support/domains/feedback"
	"github.com/turistikrota/service.support/domains/support"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Validator   *validation.Validator
	MongoDB     *mongo.DB
	CacheSrv    cache.Service
}

func NewApplication(cnf Config) app.Application {
	contactFactory := contact.NewFactory()
	contactRepo := contact.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Contact.Collection), contactFactory)

	feedbackFactory := feedback.NewFactory()
	feedbackRepo := feedback.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Feedback.Collection), feedbackFactory)

	supportFactory := support.NewFactory()
	supportRepo := support.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Support.Collection), supportFactory)

	return app.Application{
		Commands: app.Commands{
			ContactCreate:         command.NewContactCreateHandler(contactFactory, contactRepo),
			ContactRead:           command.NewContactReadHandler(contactFactory, contactRepo),
			FeedbackCreate:        command.NewFeedbackCreateHandler(feedbackFactory, feedbackRepo),
			FeedbackRead:          command.NewFeedbackReadHandler(feedbackFactory, feedbackRepo),
			SupportAdminAddMsg:    command.NewSupportAdminAddMsgHandler(supportRepo),
			SupportAdminClose:     command.NewSupportAdminCloseHandler(supportRepo),
			SupportAdminDelete:    command.NewSupportAdminDeleteHandler(supportRepo),
			SupportAdminRemoveMsg: command.NewSupportAdminRemoveMsgHandler(supportRepo),
			SupportAdminUpdate:    command.NewSupportAdminUpdateHandler(supportRepo),
			SupportCreate:         command.NewSupportCreateHandler(supportFactory, supportRepo),
			SupportAddMsg:         command.NewSupportAddMsgHandler(supportRepo),
			SupportClose:          command.NewSupportCloseHandler(supportRepo),
			SupportDelete:         command.NewSupportDeleteHandler(supportRepo),
		},
		Queries: app.Queries{
			ContactList:        query.NewContactListHandler(contactRepo),
			FeedbackList:       query.NewFeedbackListHandler(feedbackRepo),
			SupportAdminFilter: query.NewSupportAdminFilterHandler(supportRepo),
			SupportAdminGet:    query.NewSupportAdminGetHandler(supportRepo),
			SupportGet:         query.NewSupportGetHandler(supportRepo),
			SupportFilter:      query.NewSupportFilterHandler(supportRepo),
		},
	}
}
