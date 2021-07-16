package util

type BaseError struct {
	Code int
	Message string
}

func NewMyError(code int, msg string) error {
	return &BaseError{Code: code, Message: msg}
}

func (t *BaseError) Error() string {
	return t.Message
}
