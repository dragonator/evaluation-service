package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dragonator/evaluation-service/module/evaluation/internal/evaluator"
)

var (
	_evaluateEndpoint = "/evaluate"
	_validateEndpoint = "/validate"

	_numberPart       = `(0|[1-9][0-9]*(\.[0-9]*)?)`
	_operationPart    = fmt.Sprintf(` +([a-z ]+[a-z])(?: +%s)?`, _numberPart)
	_isMathRegex      = regexp.MustCompile(fmt.Sprintf(`^What +is +%s`, _numberPart))
	_syntaxCheckRegex = regexp.MustCompile(fmt.Sprintf(`^(?:%s)* *\?`, _operationPart))
	_operationRegex   = regexp.MustCompile(_operationPart)
)

// Parser implements functionality for parsing an expression string to expression object.
type Parser struct{}

// NewParser is a construction function for Parser.
func NewParser() *Parser {
	return &Parser{}
}

// ParseOperation parses single operation string to Operation object.
func (p *Parser) ParseOperation(operation, arg string) (evaluator.Operation, error) {
	var op evaluator.Operation

	operationPrefix := strings.Split(operation, " ")[0]

	var opType evaluator.OperationType

	switch {
	case strings.HasPrefix(evaluator.PlusOp.String(), operationPrefix):
		opType = evaluator.PlusOp

	case strings.HasPrefix(evaluator.MinusOp.String(), operationPrefix):
		opType = evaluator.MinusOp

	case strings.HasPrefix(evaluator.DivisionOp.String(), operationPrefix):
		opType = evaluator.DivisionOp

	case strings.HasPrefix(evaluator.MultiplicationOp.String(), operationPrefix):
		opType = evaluator.MultiplicationOp

	default:
		return op, fmt.Errorf("%w: %s", ErrUnsupportedOperation, operation)
	}

	if operation != opType.String() {
		return op, fmt.Errorf("%w: %s", ErrInvalidOperationSyntax, operation)
	}

	opArg, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return op, fmt.Errorf("converting to integer: %v", err)
	}

	op.Type = opType
	op.Arg = opArg

	return op, nil
}

// ParseOperation parses single expression string to Expression object.
func (p *Parser) ParseExpression(expression string) (evaluator.Expression, error) {
	var expr evaluator.Expression

	isMathResult := _isMathRegex.FindStringSubmatch(expression)
	if len(isMathResult) == 0 {
		return expr, ErrNonMathQuestion
	}

	expression = strings.Replace(expression, isMathResult[0], "", 1)

	if !_syntaxCheckRegex.MatchString(expression) {
		return expr, ErrInvalidSyntax
	}

	startingValue, err := strconv.ParseFloat(isMathResult[1], 64)
	if err != nil {
		return expr, fmt.Errorf("converting to integer: %v", err)
	}

	expr.StartingValue = startingValue

	operations := _operationRegex.FindAllStringSubmatch(expression, -1)
	for _, o := range operations {
		op, err := p.ParseOperation(o[1], o[2])
		if err != nil {
			return expr, err
		}

		expr.Operations = append(expr.Operations, op)
	}

	return expr, nil
}
