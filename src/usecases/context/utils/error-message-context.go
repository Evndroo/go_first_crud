package utils

import (
	"context"
)

type ErrorMessages struct {
	InternalServerError string
	InvalidIdBadRequest string
}

func WithErrorMessagesContext(ctx context.Context) context.Context {

	errorMessages := ErrorMessages{
		InternalServerError: "Sorry, we have a problem, please try again later.",
		InvalidIdBadRequest: "Please send a valid id.",
	}

	return context.WithValue(ctx, "errorMessages", errorMessages)
}

func GetErrorMessagesFromContext(ctx context.Context) (ErrorMessages, bool) {
	errorMessages, success := ctx.Value("errorMessages").(ErrorMessages)

	if !success {
		return ErrorMessages{}, false
	}

	return errorMessages, true
}
