package support

import (
	"time"

	"github.com/google/uuid"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newSupportErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewConfig struct {
	User    *User
	Subject string
	Message string
}

func (f Factory) New(cnf NewConfig) *Entity {
	t := time.Now()
	firstMessage := &Message{
		UUID:    uuid.New().String(),
		Text:    cnf.Message,
		IsAdmin: false,
		Date:    t,
	}
	return &Entity{
		User:    cnf.User,
		Subject: cnf.Subject,
		Messages: []*Message{
			firstMessage,
		},
		State:        States.Open,
		IsUserClosed: false,
		ClosedAt:     nil,
		UpdatedAt:    &t,
		CreatedAt:    t,
	}
}
