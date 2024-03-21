package errorutil

import "errors"

// Error list.
var (
	// General errors
	ErrInvalidDBFormat      = errors.New("invalid db address")
	ErrNotFoundBoilerplate  = errors.New("not found boilerplate")
	ErrInvalidRequestFormat = errors.New("invalid request format")
	ErrInternalDB           = errors.New("internal database error")
	ErrInternalElastic      = errors.New("internal elastic error")
	ErrInternalCache        = errors.New("internal cache error")
	ErrInternalServer       = errors.New("internal server error")

	// Products errors
	ErrNotFoundProduct = errors.New("not found product")
)
