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
	ContactCreate         command.ContactCreateHandler
	ContactRead           command.ContactReadHandler
	FeedbackCreate        command.FeedbackCreateHandler
	FeedbackRead          command.FeedbackReadHandler
	SupportAdminAddMsg    command.SupportAdminAddMsgHandler
	SupportAdminClose     command.SupportAdminCloseHandler
	SupportAdminDelete    command.SupportAdminDeleteHandler
	SupportAdminRemoveMsg command.SupportAdminRemoveMsgHandler
	SupportAdminUpdate    command.SupportAdminUpdateHandler
	SupportCreate         command.SupportCreateHandler
	SupportAddMsg         command.SupportAddMsgHandler
	SupportClose          command.SupportCloseHandler
	SupportDelete         command.SupportDeleteHandler
}

type Queries struct {
	ContactList        query.ContactListHandler
	FeedbackList       query.FeedbackListHandler
	SupportAdminFilter query.SupportAdminFilterHandler
	SupportAdminGet    query.SupportAdminGetHandler
	SupportGet         query.SupportGetHandler
	SupportFilter      query.SupportFilterHandler
}
