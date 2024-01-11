package support

type messages struct {
	Failed      string
	InvalidUUID string
	NotFound    string
}

var i18nMessages = messages{
	Failed:      "support_failed",
	InvalidUUID: "support_invalid_uuid",
	NotFound:    "support_not_found",
}
