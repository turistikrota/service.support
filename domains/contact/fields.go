package contact

type fieldsType struct {
	UUID    string
	Subject string
	Email   string
	Message string
	IsRead  string
	Date    string
}

var fields = fieldsType{
	UUID:    "_id",
	Subject: "subject",
	Email:   "email",
	Message: "message",
	IsRead:  "is_read",
	Date:    "date",
}
