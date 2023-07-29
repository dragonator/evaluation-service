package evaluator

import "fmt"

// OperationType represents an operation type.
type OperationType string

// String returns the operation type as string.
func (ot OperationType) String() string {
	return string(ot)
}

// Operation definitions.
const (
	PlusOp           OperationType = "plus"
	MinusOp          OperationType = "minus"
	DivisionOp       OperationType = "divided by"
	MultiplicationOp OperationType = "multiplied by"
)

// Operation represents a single operation.
type Operation struct {
	Type OperationType
	Arg  float64
}

// Expression represents a single expression.
type Expression struct {
	StartingValue float64
	Operations    []Operation
}

// Evaluator implements functionality for evaluating an expression object.
type Evaluator struct{}

// NewEvaluator is a construction function for Evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

// Evaluate evaluates a single Expression object.
func (e *Evaluator) Evaluate(expr Expression) (float64, error) {
	res := expr.StartingValue

	for _, o := range expr.Operations {
		switch o.Type {
		case PlusOp:
			res += o.Arg
		case MinusOp:
			res -= o.Arg
		case DivisionOp:
			res /= o.Arg
		case MultiplicationOp:
			res *= o.Arg
		default:
			return 0, fmt.Errorf("usupportes operation: %s", o.Type)
		}
	}

	return res, nil
}
