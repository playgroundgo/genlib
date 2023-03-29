package errors

import (
	"errors"
)

// ErrEmpty signals that a collection is empty when it wasn't supposed to be.
var ErrEmpty = errors.New("container is empty")
