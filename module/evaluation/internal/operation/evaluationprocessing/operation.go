package evaluationprocessing

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/dragonator/evaluation-service/module/evaluation/internal/evaluator"
	"github.com/dragonator/evaluation-service/module/evaluation/internal/http/service/svc"
	"github.com/dragonator/evaluation-service/module/evaluation/internal/parser"
)

var (
	_evaluateEndpoint = "/evaluate"
	_validateEndpoint = "/validate"
)

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
type Operation struct {
	expressionErrors sync.Map
	evaluator        *evaluator.Evaluator
	parser           *parser.Parser
}

// NewOperation is a contruction function for Operation.
func NewOperation(parser *parser.Parser, evaluator *evaluator.Evaluator) *Operation {
	return &Operation{
		evaluator: evaluator,
		parser:    parser,
	}
}

// Evaluate parses and evaluates an expression represented by string.
func (o *Operation) Evaluate(expression string) (float64, error) {
	expr, err := o.parser.ParseExpression(expression)
	if err != nil {
		errSave := o.saveError(err, expression, _evaluateEndpoint)
		if errSave != nil {
			return 0, errSave
		}

		var e parser.ParserError
		if errors.As(err, &e) {
			return 0, &svc.Error{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}
		}

		return 0, err
	}

	res, err := o.evaluator.Evaluate(expr)
	if err != nil {
		return 0, err
	}

	return res, nil
}

// Validate parses and validates an expression represented by string.
func (o *Operation) Validate(expression string) (*ValidationResult, error) {
	var res ValidationResult

	_, err := o.parser.ParseExpression(expression)
	if err != nil {
		res.Valid = false
		res.Reason = err.Error()

		err = o.saveError(err, expression, _validateEndpoint)
		if err != nil {
			return nil, err
		}

		return &res, nil
	}

	res.Valid = true

	return &res, nil
}

// Errors returns a list of expression errors occured during evaluation or validation of expressions.
func (o *Operation) Errors() []*ExpressionError {
	res := make([]*ExpressionError, 0, 10)

	o.expressionErrors.Range(func(key, value any) bool {
		res = append(res, value.(*ExpressionError))

		return true
	})

	return res
}

func (o *Operation) saveError(err error, expression string, endpoint string) error {
	var ee *ExpressionError

	exprKey := fmt.Sprintf("%s: %s", endpoint, expression)

	v, ok := o.expressionErrors.Load(exprKey)
	if ok {
		ee = v.(*ExpressionError)
	} else {
		var e parser.ParserError
		if !errors.As(err, &e) {
			return err
		}

		ee = &ExpressionError{
			Expression: expression,
			Endpoint:   endpoint,
			Frequency:  0,
			Type:       string(e.Type),
		}
	}

	ee.Frequency++

	o.expressionErrors.Store(exprKey, ee)

	return nil
}
