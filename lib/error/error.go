package error

import (
	"errors"
	"fmt"
)

var (
	ErrRedisSetText      = errors.New("redis failed to set text")
	ErrRedisGetText      = errors.New("redis failed to get text")
	ErrUseCaseCreateText = errors.New("usecase failed to create text")
	ErrUseCaseGetText    = errors.New("usecase failed to get text")
	ErrNotFoundInput     = errors.New("not found input")
	ErrRedisNotFound     = errors.New("data not found")
)

type Error struct {
	Str string
}

type WrapperError struct {
	Context string
	Err     error
}

func (err *WrapperError) Error() string {
	return fmt.Sprintf("%s: %v", err.Context, err.Err)
}

func NewError(ctx string, err error) error {

	return &WrapperError{
		Context: ctx,
		Err:     err,
	}
}

type RestfulError struct {
	Err string `json:"error"`
}
