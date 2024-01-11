package support

type AdminListDto struct{}

type ListDto struct{}

type AdminDto struct{}

type DetailDto struct{}

func (e *Entity) ToList() ListDto {
	return ListDto{}
}

func (e *Entity) ToAdmin() AdminDto {
	return AdminDto{}
}

func (e *Entity) ToDetail() DetailDto {
	return DetailDto{}
}

func (e *Entity) ToAdminList() AdminListDto {
	return AdminListDto{}
}
