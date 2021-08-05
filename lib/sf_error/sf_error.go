package sf_error

import (
	"fmt"
	"runtime"
)

type SfErrorCode int

const (
	UnknownError = iota
	NotFoundError
	ParameterError
	InternalError
)

type SfError struct {
	Code     SfErrorCode
	Msg      string
	Function string
	File     string
	Line     int
	Err      error
}

func NewSfError(code SfErrorCode, msg string) *SfError {
	err := SfError{
		Code: code,
		Msg:  msg,
	}
	pc, fn, line, ok := runtime.Caller(1)
	if ok {
		err.Function = runtime.FuncForPC(pc).Name()
		err.File = fn
		err.Line = line
	}
	return &err
}

func (err *SfError) Wrap(message string) *SfError {
	pc, fn, line, _ := runtime.Caller(1)
	s := fmt.Sprintf("[%s] %s:%d %s\n", message, fn, line, runtime.FuncForPC(pc).Name())
	err.Err = fmt.Errorf("%w %s", err.Err, s)
	return err
}

func (err *SfError) Error() string {
	switch err.Code {
	case NotFoundError:
		return fmt.Sprintf("ERROR: NotFoundError %s", err.Msg)
	case ParameterError:
		return fmt.Sprintf("ERROR: ParameterError %s", err.Msg)
	case InternalError:
		return fmt.Sprintf("ERROR: InternalError %s", err.Msg)
	}
	return fmt.Sprintf("ERROR: UnknownError")
}