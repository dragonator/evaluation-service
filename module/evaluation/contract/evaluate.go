package contract

// EvaluateRequest is a client request object for evaluating expressions.
type EvaluateRequest struct {
	Expression string `json:"expression"`
}

// EvaluateResponse is a server response object for evaluating expressions.
type EvaluateResponse struct {
	Result float64 `json:"result"`
}
