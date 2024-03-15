package util

import "server/infrastructure"

type BadRequest struct {
	Message string
}

func (b BadRequest) Error() string {
	return b.Message
}

// type BodyValidationError struct {
// 	Message   string
// 	ListError any
// }

// func (validationError BodyValidationError) Error() string {
// 	return validationError.Message
// }

type Duplicate struct {
	Message string
}

func (b Duplicate) Error() string {
	return b.Message
}

type NoData struct {
	Message string
}

func (b NoData) Error() string {
	return b.Message
}

type CustomBadRequest struct {
	temp        string
	Errors      any
	Messages    string
	StatusCodes int
}

func (validationError CustomBadRequest) Error() string {
	return validationError.temp
}

func (validationError CustomBadRequest) CustomError() any {
	return validationError.Errors
}

func (validationError CustomBadRequest) Message() string {
	if validationError.Messages == "" {
		validationError.Messages = infrastructure.Localize("BAD_REQUEST")
	}
	return validationError.Messages
}

func (validationError CustomBadRequest) StatusCode() int {
	if validationError.StatusCodes == 0 {
		validationError.StatusCodes = 400
	}
	return validationError.StatusCodes
}
