package app

import (
	"github.com/turistikrota/service.support/app/command"
	"github.com/turistikrota/service.support/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ContactCreate  command.ContactCreateHandler
	ContactRead    command.ContactReadHandler
	FeedbackCreate command.FeedbackCreateHandler
	FeedbackRead   command.FeedbackReadHandler
}

type Queries struct {
	ContactList  query.ContactListHandler
	FeedbackList query.FeedbackListHandler
}
