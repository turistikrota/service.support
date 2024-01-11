package support

import "time"

type AdminListDto struct {
	UUID      string     `json:"uuid"`
	User      User       `json:"user"`
	Subject   string     `json:"subject"`
	State     State      `json:"state"`
	UpdatedAt *time.Time `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
}

type ListDto struct {
	UUID      string     `json:"uuid"`
	Subject   string     `json:"subject"`
	State     State      `json:"state"`
	UpdatedAt *time.Time `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
}

type AdminDto struct {
	UUID         string     `json:"uuid"`
	User         User       `json:"user"`
	Subject      string     `json:"subject"`
	Messages     []*Message `json:"messages"`
	State        State      `json:"state"`
	IsUserClosed bool       `json:"isUserClosed"`
	ClosedAt     *time.Time `json:"closedAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	CreatedAt    time.Time  `json:"createdAt"`
}

type UserMessageDto struct {
	Text    string    `json:"text"`
	IsAdmin bool      `json:"isAdmin"`
	Date    time.Time `json:"date"`
}

type DetailDto struct {
	UUID         string            `json:"uuid"`
	Subject      string            `json:"subject"`
	Messages     []*UserMessageDto `json:"messages"`
	State        State             `json:"state"`
	IsUserClosed bool              `json:"isUserClosed"`
	ClosedAt     *time.Time        `json:"closedAt"`
	UpdatedAt    *time.Time        `json:"updatedAt"`
}

func (e *Entity) ToList() ListDto {
	return ListDto{
		UUID:      e.UUID,
		Subject:   e.Subject,
		State:     e.State,
		UpdatedAt: e.UpdatedAt,
		CreatedAt: e.CreatedAt,
	}
}

func (e *Entity) ToAdmin() AdminDto {
	return AdminDto{
		UUID:         e.UUID,
		User:         *e.User,
		Subject:      e.Subject,
		Messages:     e.Messages,
		State:        e.State,
		IsUserClosed: e.IsUserClosed,
		ClosedAt:     e.ClosedAt,
		UpdatedAt:    e.UpdatedAt,
		CreatedAt:    e.CreatedAt,
	}
}

func (e *Entity) ToDetail() DetailDto {
	messages := make([]*UserMessageDto, 0)
	for _, message := range e.Messages {
		messages = append(messages, &UserMessageDto{
			Text:    message.Text,
			IsAdmin: message.IsAdmin,
			Date:    message.Date,
		})
	}
	return DetailDto{
		UUID:         e.UUID,
		Subject:      e.Subject,
		Messages:     messages,
		State:        e.State,
		IsUserClosed: e.IsUserClosed,
		ClosedAt:     e.ClosedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}

func (e *Entity) ToAdminList() AdminListDto {
	return AdminListDto{
		UUID:      e.UUID,
		User:      *e.User,
		Subject:   e.Subject,
		State:     e.State,
		UpdatedAt: e.UpdatedAt,
		CreatedAt: e.CreatedAt,
	}
}
