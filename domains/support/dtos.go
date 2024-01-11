package support

type AdminListDto struct{}

type ListDto struct{}

func (e *Entity) ToList() ListDto {
	return ListDto{}
}

func (e *Entity) ToAdminList() AdminListDto {
	return AdminListDto{}
}
