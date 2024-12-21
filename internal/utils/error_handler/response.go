package errorhandler

import (
	"go.uber.org/zap"
)

type APIError struct {
	Error string `json:"error"`
}

func NewAPIError(err error, defaultStatus ...int) (code int, i interface{}) {
	zap.L().Error("error", zap.Error(err))
	response := &APIError{
		Error: err.Error(),
	}

	status := 500

	if len(defaultStatus) > 0 {
		status = defaultStatus[0]
	} else {
		status = getStatusCode(err)
	}
	return status, response
}

func getStatusCode(err error) int {
	if code, ok := ERRORS_MAP[err]; ok {
		return code
	}

	return 500
}
