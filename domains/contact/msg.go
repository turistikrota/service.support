package contact

type messages struct {
	Failed      string
	InvalidUUID string
}

var i18nMessages = messages{
	Failed:      "contact_failed",
	InvalidUUID: "contact_invalid_uuid",
}
