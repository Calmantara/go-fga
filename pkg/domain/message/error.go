package message

type ErrorType string

type ErrorMessage struct {
	Error error     `json:"error"`
	Type  ErrorType `json:"error_type"`
}
