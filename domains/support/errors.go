package support

import "github.com/cilloparch/cillop/i18np"

type Errors interface {
	Failed(string) *i18np.Error
	InvalidUUID() *i18np.Error
}

type supportErrors struct{}

func newSupportErrors() Errors {
	return &supportErrors{}
}

func (e *supportErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(i18nMessages.Failed, i18np.P{
		"Operation": operation,
	})
}

func (e *supportErrors) InvalidUUID() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidUUID)
}
