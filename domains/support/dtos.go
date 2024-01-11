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

type DetailDto struct {
	UUID         string     `json:"uuid"`
	Subject      string     `json:"subject"`
	Messages     []*Message `json:"messages"`
	State        State      `json:"state"`
	IsUserClosed bool       `json:"isUserClosed"`
	ClosedAt     *time.Time `json:"closedAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
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
	return DetailDto{
		UUID:         e.UUID,
		Subject:      e.Subject,
		Messages:     e.Messages,
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
