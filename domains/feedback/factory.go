package feedback

import "time"

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newFeedbackErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewConfig struct {
	OS      string `json:"os" bson:"os"`
	Version string `json:"version" bson:"version"`
	User    *User  `json:"user" bson:"user"`
	Message string `json:"message" bson:"message"`
}

func (f Factory) New(cnf NewConfig) *Entity {
	read := false
	return &Entity{
		OS:      cnf.OS,
		Version: cnf.Version,
		User:    cnf.User,
		Message: cnf.Message,
		IsRead:  &read,
		Date:    time.Now(),
	}
}
