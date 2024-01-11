package support

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
