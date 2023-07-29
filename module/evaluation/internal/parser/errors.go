package parser

import (
	"errors"
)

type ExpressionErrorType string

const (
	ErrTypeInvalidSyntax        = "ErrInvalidSyntax"
	ErrTypeUnsupportedOperation = "ErrUnsupportedOperation"
	ErrTypeNonMathQuestion      = "ErrNonMathQuestion"
)

type ParserError struct {
	Type ExpressionErrorType
	Err  error
}

func (ee ParserError) Error() string {
	return ee.Err.Error()
}

var (
	ErrInvalidSyntax        = ParserError{Type: ErrTypeInvalidSyntax, Err: errors.New("invalid expression syntax")}
	ErrUnsupportedOperation = ParserError{Type: ErrTypeUnsupportedOperation, Err: errors.New("unsupported operation")}
	ErrNonMathQuestion      = ParserError{Type: ErrTypeNonMathQuestion, Err: errors.New("non-math question")}
)
