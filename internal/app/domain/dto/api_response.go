package dto

type APIResponse[T any, D any] struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorMessage T      `json:"error_message"`
	Data         D      `json:"data"`
}
