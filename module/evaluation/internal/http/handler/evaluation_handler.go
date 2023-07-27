package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dragonator/evaluation-service/module/evaluation/contract"
	"github.com/dragonator/evaluation-service/module/evaluation/internal/http/service/svc"
	"github.com/dragonator/evaluation-service/module/evaluation/internal/operation/evaluationprocessing"
)

// EvaluationProcessingOp is a contract to a evaluation processing operation.
type EvaluationProcessingOp interface {
	Evaluate(ctx context.Context, expression string) (float64, error)
	Validate(ctx context.Context, expression string) (*evaluationprocessing.ValidationResult, error)
	Errors(ctx context.Context) ([]*evaluationprocessing.ExpressionError, error)
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
	return func(w http.ResponseWriter, r *http.Request) {
		req := &contract.EvaluateRequest{}
		if err := decode(r, req); err != nil {
			errorResponse(w, fmt.Errorf("%w: %v", svc.ErrDecodeRequest, err))
			return
		}

		res, err := rh.evaluationProcessingOp.Evaluate(r.Context(), req.Expression)
		if err != nil {
			errorResponse(w, err)
			return
		}

		successResponse(w, &contract.EvaluateResponse{
			Result: res,
		})

		return
	}
}

// Validate validates the expressions sent as request and returns whether they are valid.
func (rh *EvaluationHandler) Validate() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &contract.ValidateRequest{}
		if err := decode(r, req); err != nil {
			errorResponse(w, fmt.Errorf("%w: %v", svc.ErrDecodeRequest, err))
			return
		}

		result, err := rh.evaluationProcessingOp.Validate(r.Context(), req.Expression)
		if err != nil {
			errorResponse(w, err)
			return
		}

		successResponse(w, &contract.ValidateResponse{
			Valid:  result.Valid,
			Reason: result.Reason,
		})

		return
	}
}

// Errors returns all errors occurred during evalating expressions.
func (rh *EvaluationHandler) Errors() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := rh.evaluationProcessingOp.Errors(r.Context())
		if err != nil {
			errorResponse(w, err)
			return
		}

		var resp contract.ErrorsResponse
		for _, e := range result {
			resp = append(resp, contract.Error{
				Expression: e.Expression,
				Endpoint:   e.Endpoint,
				Frequency:  e.Frequency,
				Type:       e.Type,
			})
		}

		successResponse(w, resp)

		return
	}
}
