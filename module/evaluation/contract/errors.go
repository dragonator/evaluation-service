package contract

// ErrorsRequest is a client request object for validating expressions.
type ErrorsRequest struct{}

// Error represents a single error in the response.
type Error struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	Type       string `json:"type"`
}

// ErrorsResponse is a server response object for validating expressions.
type ErrorsResponse []Error
