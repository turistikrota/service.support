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

	return app.Application{
		Commands: app.Commands{
			ContactCreate:  command.NewContactCreateHandler(contactFactory, contactRepo),
			ContactRead:    command.NewContactReadHandler(contactFactory, contactRepo),
			FeedbackCreate: command.NewFeedbackCreateHandler(feedbackFactory, feedbackRepo),
			FeedbackRead:   command.NewFeedbackReadHandler(feedbackFactory, feedbackRepo),
		},
		Queries: app.Queries{
			ContactList:  query.NewContactListHandler(contactRepo),
			FeedbackList: query.NewFeedbackListHandler(feedbackRepo),
		},
	}
}
