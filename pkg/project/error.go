package project

type errorCode int

const (
	errExceedPoolCapacity errorCode = iota
	errNotFound
	errTimeout
	errOtherFailure
)

// internError is the error type used by all subcomponents
type internError struct {
	code  errorCode
	cause error
	ctx   []interface{}
}

func (e internError) Error() string {
	return e.cause.Error()
}

// case internalError to http response
func (e internError) HTTPResponse() error {

	return nil
}
