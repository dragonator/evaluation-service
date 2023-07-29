package contract

// ValidateRequest is a client request object for validating expressions.
type ValidateRequest struct {
	Expression string `json:"expression,omitempty"`
}

// ValidateResponse is a server response object for validating expressions.
type ValidateResponse struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}
