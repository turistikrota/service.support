package contact

import "time"

type Entity struct {
	UUID    string    `json:"uuid" bson:"_id,omitempty"`
	Subject string    `json:"subject" bson:"subject"`
	Email   string    `json:"email" bson:"email"`
	Message string    `json:"message" bson:"message"`
	IsRead  *bool     `json:"is_read" bson:"is_read"`
	Date    time.Time `json:"date" bson:"date"`
}
