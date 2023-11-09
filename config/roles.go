package config

import "github.com/turistikrota/service.shared/base_roles"

type supportRoles struct {
	ReadContact  string
	ListContact  string
	ReadFeedback string
	ListFeedback string
}

type roles struct {
	base_roles.Roles
	Support supportRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Support: supportRoles{
		ReadContact:  "support.contact.read",
		ListContact:  "support.contact.list",
		ReadFeedback: "support.feedback.read",
		ListFeedback: "support.feedback.list",
	},
}
