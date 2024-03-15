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

var statusMessages = map[int]string{
	200: "OK",
	201: "Created",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Payload Too Large",
	414: "URI Too Long",
	415: "Unsupported Media Type",
	416: "Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	421: "Misdirected Request",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Too Early",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
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
	return r.fiberCtx.Status(code).JSON(SuccessResponse{
		Data:    data,
		Meta:    meta,
		Message: message,
		Response: Response{
			Code:    code,
			Message: statusMessages[code],
		},
	})
}

func (r *responseStruct) Error(errors any, message string, statusCode int) error {
	responseMessage, exists := statusMessages[statusCode]
	if !exists {
		responseMessage = "Bad Gateway"
	}

	return r.fiberCtx.Status(statusCode).JSON(ErrorResponse{
		Error:   errors,
		Message: message,
		Response: Response{
			Code:    statusCode,
			Message: responseMessage,
		},
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
