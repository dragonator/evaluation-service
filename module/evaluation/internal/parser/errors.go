package parser

import (
	"errors"
)

// ExpressionErrorType represents a single expression error type.
type ExpressionErrorType string

// Expression error type definitions.
const (
	ErrTypeInvalidSyntax          = "ErrInvalidSyntax"
	ErrTypeInvalidOperationSyntax = "ErrInvalidOperationSyntax"
	ErrTypeUnsupportedOperation   = "ErrUnsupportedOperation"
	ErrTypeNonMathQuestion        = "ErrNonMathQuestion"
)

// ParserError represents an error occured during expression parsing.
type ParserError struct {
	Type ExpressionErrorType
	Err  error
}

// Error returns the text of the parser error.
func (ee ParserError) Error() string {
	return ee.Err.Error()
}

// Parser error definitions.
var (
	ErrInvalidSyntax          = ParserError{Type: ErrTypeInvalidSyntax, Err: errors.New("invalid expression syntax")}
	ErrInvalidOperationSyntax = ParserError{Type: ErrTypeInvalidOperationSyntax, Err: errors.New("invalid operation syntax")}
	ErrUnsupportedOperation   = ParserError{Type: ErrTypeUnsupportedOperation, Err: errors.New("unsupported operation")}
	ErrNonMathQuestion        = ParserError{Type: ErrTypeNonMathQuestion, Err: errors.New("non-math question")}
)
