package store

import (
	"errors"
)

var (
	// Returned when a player is not found in a store
	ErrPlayerNotFound = errors.New("player not found")
)
