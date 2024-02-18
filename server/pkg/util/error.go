package util

type BadRequest struct {
	Message string
}

func (b BadRequest) Error() string {
	return b.Message
}

type BodyValidationError struct {
	Message   string
	ListError any
}

func (validationError BodyValidationError) Error() string {
	return validationError.Message
}

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
