package feedback

type fieldsType struct {
	UUID    string
	OS      string
	Version string
	User    string
	Message string
	IsRead  string
	Date    string
}

type userFieldsType struct {
	UUID  string
	Name  string
	Email string
}

var fields = fieldsType{
	UUID:    "_id",
	OS:      "os",
	Version: "version",
	User:    "user",
	Message: "message",
	IsRead:  "is_read",
	Date:    "date",
}

var userFields = userFieldsType{
	UUID:  "uuid",
	Name:  "name",
	Email: "email",
}

func userField(field string) string {
	return fields.User + "." + field
}
