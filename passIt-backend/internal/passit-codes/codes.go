package codes

// Package passitcodes defines constants for various response codes used in the PassIt application.

const (
	UserCreatedSuccessfully   = 201
	UserUpdatedSuccessfully   = 200
	UserLoggedInSuccessfully  = 202
	JobsRetrievedSuccessfully = 205

	// Error codes
	GetJobBadRequest = 400
	JobIdNotFound    = 405
)
