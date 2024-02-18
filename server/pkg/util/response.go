package util

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Meta     interface{} `json:"meta,omitempty"`
	Response Response    `json:"response"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"per_page"`
	TotalRows  int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ErrorResponse struct {
	Error    any      `json:"errors"`
	Message  string   `json:"message"`
	Response Response `json:"response"`
}

type ImportResponse struct {
	Errors     []ImportError `json:"errors"`
	TotalInput int           `json:"total_input"`
	Success    int           `json:"success"`
	Failed     int           `json:"failed"`
}

type ImportError struct {
	Row    int `json:"nomor"`
	Errors any `json:"error"`
}

const (
	MESSAGE_OK_CREATE              = "Successfully inserted data."
	MESSAGE_OK_DELETE              = "Successfully deleted data."
	MESSAGE_OK_UPDATE              = "Successfully updated data."
	MESSAGE_OK_READ                = "Successfully displayed data."
	MESSAGE_OK_IMPORT              = "Successfully imported data."
	MESSAGE_OK_EXPORT              = "Successfully exported data."
	MESSAGE_OK_UPLOAD              = "Successfully uploaded data."
	MESSAGE_OK_LOGIN               = "Successfully logged in."
	MESSAGE_OK_REFRESH             = "Successfully refreshed token."
	MESSAGE_OK_LOGOUT              = "Successfully logged out."
	MESSAGE_OK_REQUEST_TOKEN_RESET = "Successfully requested token."
	MESSAGE_OK_CHANGE_PASSWORD     = "Successfully updated password."
	MESSAGE_FAILED_CREATE          = "Failed to insert data."
	MESSAGE_FAILED_DELETE          = "Failed to delete data."
	MESSAGE_FAILED_UPDATE          = "Failed to update data."
	MESSAGE_FAILED_READ            = "Failed to display data."
	MESSAGE_FAILED_IMPORT          = "Failed to import data."
	MESSAGE_FAILED_EXPORT          = "Failed to export data."
	MESSAGE_FAILED_VALIDATION      = "Failed to validate data."
	MESSAGE_BAD_REQUEST            = "Bad request."
	MESSAGE_BAD_SYSTEM             = "Server error."
	MESSAGE_UNAUTHORIZED           = "Unauthorized."
	MESSAGE_CONFLICT               = "Data already exists."
	MESSAGE_NOT_FOUND              = "Data not found."
	MESSAGE_FAILED_LOGIN           = "Incorrect username or password."
)

var statusMessages = map[int]string{
	200: "OK",
	201: "Created",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	409: "Conflict",
	415: "Unsupported Media Type",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
}

type repsonseInterface interface {
	Success(data any, meta any, message string, statusCode ...int) error
	Error(errors any, message string, statusCode int) error
	Import(errors []ImportError, totalInput int, failed int) error
}

type responseStruct struct {
	fiberCtx *fiber.Ctx
}

func NewResponse(fiberCtx *fiber.Ctx) repsonseInterface {
	return &responseStruct{
		fiberCtx: fiberCtx,
	}
}

func (r *responseStruct) Success(data any, meta any, message string, statusCode ...int) error {
	code := 200
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	response := Response{
		Code:    code,
		Message: statusMessages[code],
	}
	return r.fiberCtx.Status(code).JSON(SuccessResponse{
		Data:     data,
		Meta:     meta,
		Message:  message,
		Response: response,
	})
}

func (r *responseStruct) Error(errors any, message string, statusCode int) error {
	responseMessage := statusMessages[statusCode]
	if responseMessage == "" {
		responseMessage = "Bad Gateway"
	}

	response := Response{
		Code:    statusCode,
		Message: responseMessage,
	}

	return r.fiberCtx.Status(statusCode).JSON(ErrorResponse{
		Error:    errors,
		Message:  message,
		Response: response,
	})
}

func (r *responseStruct) Import(errors []ImportError, totalInput int, failed int) error {
	return r.fiberCtx.Status(fiber.StatusOK).JSON(ImportResponse{
		Errors:     errors,
		TotalInput: totalInput,
		Success:    totalInput - failed,
		Failed:     failed,
	})
}
