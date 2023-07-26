package handler

import (
	"context"
	"net/http"

	"github.com/dragonator/evaluation-service/module/evaluation/internal/operation/evaluationprocessing"
)

// EvaluationProcessingOp is a contract to a evaluation processing operation.
type EvaluationProcessingOp interface {
	Evaluate(ctx context.Context, expressions string) (float64, error)
	Validate(ctx context.Context, expressions string) (bool, error)
	Errors(ctx context.Context, expressions string) ([]*evaluationprocessing.EvaluationError, error)
}

// EvaluationHandler holds implementation of handlers for evaluations.
type EvaluationHandler struct {
	evaluationProcessingOp EvaluationProcessingOp
}

// NewEvaluationHandler is a construction function for EvaluationHandler.
func NewEvaluationHandler(evaluationProcessingOp EvaluationProcessingOp) *EvaluationHandler {
	return &EvaluationHandler{
		evaluationProcessingOp: evaluationProcessingOp,
	}
}

// Evaluate evaluates the expressions send as request and returns the result as response.
func (rh *EvaluationHandler) Evaluate() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// Validate validates the expressions sent as request and returns whether they are valid.
func (rh *EvaluationHandler) Validate() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// Errors returns all errors occurred during evalating expressions.
func (rh *EvaluationHandler) Errors() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
