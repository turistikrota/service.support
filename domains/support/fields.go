package support

type fieldsType struct {
	UUID         string
	User         string
	Interests    string
	Subject      string
	Messages     string
	State        string
	IsUserClosed string
	ClosedAt     string
	UpdatedAt    string
	CreatedAt    string
}

type userFieldsType struct {
	UUID string
	Name string
}

type interestFieldsType struct {
	UUID  string
	Email string
}

type messageFieldsType struct {
	UUID         string
	InterestUUID string
	Text         string
	IsAdmin      string
	IsDeleted    string
	Date         string
}

var fields = fieldsType{
	UUID:         "_id",
	User:         "user",
	Interests:    "interests",
	Subject:      "subject",
	Messages:     "messages",
	State:        "state",
	IsUserClosed: "is_user_closed",
	ClosedAt:     "closed_at",
	UpdatedAt:    "updated_at",
	CreatedAt:    "created_at",
}

var userFields = userFieldsType{
	UUID: "uuid",
	Name: "name",
}

var interestFields = interestFieldsType{
	UUID:  "uuid",
	Email: "email",
}

var messageFields = messageFieldsType{
	UUID:         "uuid",
	InterestUUID: "interest_uuid",
	Text:         "text",
	IsAdmin:      "is_admin",
	IsDeleted:    "is_deleted",
	Date:         "date",
}

func userField(field string) string {
	return fields.User + "." + field
}

func interestField(field string) string {
	return fields.Interests + "." + field
}

func messageField(field string) string {
	return fields.Messages + "." + field
}
