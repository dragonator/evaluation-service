package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dragonator/evaluation-service/module/evaluation/contract"
	"github.com/dragonator/evaluation-service/module/evaluation/internal/http/service/svc"
)

const (
	_contentTypeHeaderName = "Content-Type"
	_contentTypeJSON       = "application/json"
	_xContentTypeOptions   = "X-Content-Type-Options"
	_noSniff               = "nosniff"
)

func errorResponse(w http.ResponseWriter, err error) {
	er := &contract.ErrorResponse{Message: err.Error()}
	w.Header().Set(_contentTypeHeaderName, _contentTypeJSON)
	w.Header().Set(_xContentTypeOptions, _noSniff)

	var e *svc.Error
	if errors.As(err, &e) {
		w.WriteHeader(e.StatusCode)
		json.NewEncoder(w).Encode(er)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	er.Message = http.StatusText(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(er)
}

func successResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set(_contentTypeHeaderName, _contentTypeJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func decode(r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	return d.Decode(v)
}
