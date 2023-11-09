package feedback

import "github.com/cilloparch/cillop/i18np"

type Errors interface {
	Failed(string) *i18np.Error
	InvalidUUID() *i18np.Error
}

type feedbackErrors struct{}

func newFeedbackErrors() Errors {
	return &feedbackErrors{}
}

func (e *feedbackErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(i18nMessages.Failed, i18np.P{
		"Operation": operation,
	})
}

func (e *feedbackErrors) InvalidUUID() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidUUID)
}
