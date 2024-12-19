package response

import (
	"fmt"
	"notify-backend/api/utils/types"
	"notify-backend/config"
)

func SuccessResponse(message string, data interface{}) types.Success {
	if message == "" {
		message = "OK"
	}
	return types.Success{
		Status:  true,
		Data:    data,
		Error:   nil,
	}
}

func ErrorResponse(message string, error interface{}) types.Error {
	if message == "" {
		message = "Error"
	} else {
		messageCode := config.GetCodeError(message)
		message = fmt.Sprintf("%s. %s", message, messageCode)
	}
	return types.Error{
		Status: false,
		Error:  message,
	}
}