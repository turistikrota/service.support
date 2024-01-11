package support

type AdminListDto struct{}

func (e *Entity) ToAdminList() AdminListDto {
	return AdminListDto{}
}
