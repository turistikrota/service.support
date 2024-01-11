package command

type SupportDetailCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}
