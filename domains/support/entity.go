package support

import "time"

type Entity struct {
	UUID         string     `json:"uuid" bson:"_id,omitempty"`
	User         *User      `json:"user" bson:"user"`
	Subject      string     `json:"subject" bson:"subject"`
	Messages     []*Message `json:"messages" bson:"messages"`
	State        State      `json:"state" bson:"state"`
	IsUserClosed bool       `json:"isUserClosed" bson:"is_user_closed"`
	ClosedAt     *time.Time `json:"closedAt" bson:"closed_at"`
	UpdatedAt    *time.Time `json:"updatedAt" bson:"updated_at"`
	CreatedAt    time.Time  `json:"createdAt" bson:"created_at"`
}

type User struct {
	UUID string `json:"uuid" bson:"uuid"`
	Name string `json:"name" bson:"name"`
}

type Interest struct {
	UUID  string `json:"uuid" bson:"uuid"`
	Email string `json:"email" bson:"email"`
}

type Message struct {
	UUID         string    `json:"uuid" bson:"uuid"`
	InterestUUID string    `json:"interest_uuid,omitempty" bson:"interest_uuid,omitempty"`
	Text         string    `json:"text" bson:"text"`
	IsAdmin      bool      `json:"is_admin" bson:"is_admin"`
	IsDeleted    bool      `json:"is_deleted" bson:"is_deleted"`
	Date         time.Time `json:"date" bson:"date"`
}

type State string

type states struct {
	Open     State
	Answered State
	Closed   State
	Deleted  State
}

var States = states{
	Open:     "open",
	Answered: "answered",
	Closed:   "closed",
	Deleted:  "deleted",
}
