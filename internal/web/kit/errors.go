package kit

import "errors"

var (
	ErrInvalidRequestParams = errors.New("invalid request parameters")
	ErrIncorrectJSON        = errors.New("json is incorrect")
)
