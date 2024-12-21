package utils

type APIResponse[Data any] struct {
	Message string `json:"message"`
	Data    *Data  `json:"data,omitempty"`
}

func NewAPIResponse[Data any](message string, data ...Data) *APIResponse[Data] {
	response := &APIResponse[Data]{
		Message: message,
	}

	if len(data) > 0 {
		response.Data = &data[0]
	}

	return response
}

func NewAPIErrorResponse(code int, error error) (int, *APIResponse[any]) {
	return code, &APIResponse[any]{Message: error.Error()}
}
