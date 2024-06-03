package utils

import "time"

type Response struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

func NewSuccessResponse(data any) Response {
	return Response{
		Code:      200,
		Success:   true,
		Data:      data,
		Timestamp: time.Now(),
	}
}

func NewErrorResponse(code int, data any) Response {
	return Response{
		Code:      code,
		Success:   false,
		Data:      data,
		Timestamp: time.Now(),
	}
}
