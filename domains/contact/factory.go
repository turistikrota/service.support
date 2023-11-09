package contact

import "time"

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
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (f Factory) New(cnf NewConfig) *Entity {
	read := false
	return &Entity{
		Subject: cnf.Subject,
		Email:   cnf.Email,
		Message: cnf.Message,
		IsRead:  &read,
		Date:    time.Now(),
	}
}
