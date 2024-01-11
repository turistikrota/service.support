package config

import "github.com/turistikrota/service.shared/base_roles"

type supportRoles struct {
	ReadContact      string
	ListContact      string
	ReadFeedback     string
	ListFeedback     string
	SupportList      string
	SupportClose     string
	SupportDelete    string
	SupportRemoveMsg string
	SupportUpdate    string
	SupportView      string
	SupportSuper     string
	SupportAddMsg    string
}

type roles struct {
	base_roles.Roles
	Support supportRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Support: supportRoles{
		ReadContact:      "support.contact.read",
		ListContact:      "support.contact.list",
		ReadFeedback:     "support.feedback.read",
		ListFeedback:     "support.feedback.list",
		SupportList:      "support.list",
		SupportClose:     "support.close",
		SupportDelete:    "support.delete",
		SupportRemoveMsg: "support.remove_msg",
		SupportUpdate:    "support.update",
		SupportView:      "support.view",
		SupportSuper:     "support.super",
		SupportAddMsg:    "support.add_msg",
	},
}
