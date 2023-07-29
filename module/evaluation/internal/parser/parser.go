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
	_operationPart    = fmt.Sprintf(` ([a-z ]+[a-z])(?: %s)?`, _numberPart)
	_isMathRegex      = regexp.MustCompile(fmt.Sprintf(`^What is %s`, _numberPart))
	_syntaxCheckRegex = regexp.MustCompile(fmt.Sprintf(`^(?:%s)*\?`, _operationPart))
	_operationRegex   = regexp.MustCompile(_operationPart)
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseOperation(operation, arg string) (evaluator.Operation, error) {
	var op evaluator.Operation

	switch evaluator.OperationType(operation) {
	case evaluator.PlusOp:
		op.Type = evaluator.PlusOp
	case evaluator.MinusOp:
		op.Type = evaluator.MinusOp
	case evaluator.DivisionOp:
		op.Type = evaluator.DivisionOp
	case evaluator.MultiplicationOp:
		op.Type = evaluator.MultiplicationOp
	default:
		return op, fmt.Errorf("%w: %s", ErrUnsupportedOperation, operation)
	}

	operationArg, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return op, fmt.Errorf("converting to integer: %v", err)
	}

	op.Arg = operationArg

	return op, nil
}

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
