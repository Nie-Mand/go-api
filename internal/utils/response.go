package utils

type APIResponse[D comparable] struct {
	Message string       `json:"message" example:"success"`
	Result  *interface{} `json:"result,omitempty"`
}

type PaginatedAPIResponse[D comparable] struct {
	Message  string `json:"message" example:"success"`
	Page     int    `json:"page" example:"1"`
	PageSize int    `json:"pageSize" example:"10"`
	Result   []D    `json:"result"`
}

func NewAPIResponse(message string, data ...interface{}) *APIResponse[interface{}] {
	response := &APIResponse[interface{}]{
		Message: message,
	}

	if len(data) > 0 {
		response.Result = &data[0]
	}

	return response
}

func NewPaginatedResponse[D comparable](message string, page int, size int, data []D) *PaginatedAPIResponse[D] {
	response := &PaginatedAPIResponse[D]{
		Message:  message,
		Page:     page,
		PageSize: size,
		Result:   data,
	}

	return response
}
