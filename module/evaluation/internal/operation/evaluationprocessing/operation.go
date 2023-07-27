package evaluationprocessing

import "context"

// ExpressionError represents a single evaluation error.
type ExpressionError struct {
	Expression string
	Endpoint   string
	Frequency  int
	Type       string
}

// ValidationResult represents a single validation error.
type ValidationResult struct {
	Valid  bool
	Reason string
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

func (o *Operation) Validate(ctx context.Context, expressions string) (*ValidationResult, error) {
	return nil, nil
}

func (o *Operation) Errors(ctx context.Context, expressions string) ([]*ExpressionError, error) {
	return nil, nil
}
