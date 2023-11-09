package http

type successMessages struct {
	Ok string
}

type errorMessages struct {
	RequiredAuth           string
	CurrentUserAccess      string
	AdminRoute             string
	RequiredAccountSelect  string
	ForbiddenAccountSelect string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		Ok: "http_success_ok",
	},
	Error: errorMessages{
		RequiredAuth:           "http_error_required_auth",
		CurrentUserAccess:      "http_error_current_user_access",
		AdminRoute:             "http_error_admin_route",
		RequiredAccountSelect:  "http_error_required_account_select",
		ForbiddenAccountSelect: "http_error_forbidden_account_select",
	},
}
