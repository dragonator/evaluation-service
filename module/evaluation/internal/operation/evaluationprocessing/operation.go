package evaluationprocessing

import "context"

// EvaluationError is representing a single expression error.
type EvaluationError struct {
	Expression string
	Endpoint   string
	Frequency  int
	Type       string
}

// Operation provides an API for evaluating expressions.
type Operation struct{}

// NewOperation is a contruction function for Operation.
func NewOperation() *Operation {
	return &Operation{}
}

func (o *Operation) Evaluate(ctx context.Context, expressions string) (float64, error) {
	return 0, nil
}

func (o *Operation) Validate(ctx context.Context, expressions string) (bool, error) {
	return true, nil
}

func (o *Operation) Errors(ctx context.Context, expressions string) ([]*EvaluationError, error) {
	return nil, nil
}
