package evaluator

import "fmt"

type OperationType string

const (
	PlusOp           OperationType = "plus"
	MinusOp          OperationType = "minus"
	DivisionOp       OperationType = "divided by"
	MultiplicationOp OperationType = "multiplied by"
)

type Operation struct {
	Type OperationType
	Arg  float64
}

type Expression struct {
	StartingValue float64
	Operations    []Operation
}

type Evaluator struct{}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

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
