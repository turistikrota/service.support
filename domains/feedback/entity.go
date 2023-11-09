package feedback

import "time"

type Entity struct {
	UUID    string    `json:"uuid" bson:"_id,omitempty"`
	OS      string    `json:"os" bson:"os"`
	Version string    `json:"version" bson:"version"`
	User    *User     `json:"user" bson:"user"`
	Message string    `json:"message" bson:"message"`
	IsRead  *bool     `json:"is_read" bson:"is_read"`
	Date    time.Time `json:"date" bson:"date"`
}

type User struct {
	UUID  string `json:"uuid" bson:"uuid"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
