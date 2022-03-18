package formatter

import "net/http"

type responseFormat struct {
	Status  int32       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	successMessage = "Success"
	failMessage    = "Fail"
)

// Standard response formatter with all user input freedom
func ResponseFormatter(status int32, message string, data interface{}) responseFormat {
	return responseFormat{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// Response with 400 status and fail message, you can pass data param with error details arguments
func BadRequestResponse(data interface{}) responseFormat {
	return responseFormat{
		Status:  http.StatusBadRequest,
		Message: failMessage,
		Data:    data,
	}
}

// Response with 404 status and fail message, you can pass data param with error details arguments
func NotFoundResponse(data interface{}) responseFormat {
	return responseFormat{
		Status:  http.StatusNotFound,
		Message: failMessage,
		Data:    data,
	}
}

// Response with 401 status and fail message, you can pass data param with error details arguments
func UnauthorizedResponse(data interface{}) responseFormat {
	return responseFormat{
		Status:  http.StatusUnauthorized,
		Message: failMessage,
		Data:    data,
	}
}

// Response with 500 status and fail message, you can pass data param with error details arguments
func InternalServerErrorResponse(data interface{}) responseFormat {
	return responseFormat{
		Status:  http.StatusInternalServerError,
		Message: failMessage,
		Data:    data,
	}
}

// Response with 200 status and fail message, you can pass data with actual data param return arguments
func SuccessResponse(data interface{}) responseFormat {
	return responseFormat{
		Status:  http.StatusOK,
		Message: successMessage,
		Data:    data,
	}
}
