package constant

import "time"

// Package constant defines application-wide constants

var (
	// SessionDuration defines the duration for which a session is valid
	SessionDuration = time.Duration(0.5 * float64(time.Hour))
)
