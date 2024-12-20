package response

import (
	"notify-backend/api/utils/types"
)

func SuccessResponse(message string, data interface{}) types.Success {
	if message == "" {
		message = "OK"
	}
	return types.Success{
		Status:  true,
		Data:    data,
		Message: message,
	}
}

func ErrorResponse(message string, error interface{}) types.Error {
	if message == "" {
		message = "Error"
	}
	return types.Error{
		Status: false,
		Error:  message,
	}
}
