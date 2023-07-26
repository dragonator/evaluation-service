package service

import (
	"net/http"

	"github.com/go-chi/chi"
)

// EvaluationHandler is a contract to a notification handler.
type EvaluationHandler interface {
	Evaluate() func(w http.ResponseWriter, r *http.Request)
	Validate() func(w http.ResponseWriter, r *http.Request)
	Errors() func(w http.ResponseWriter, r *http.Request)
}

// NewRouter is a construction function for router that handles evaluation operations.
func NewRouter(nh EvaluationHandler) http.Handler {
	router := chi.NewRouter()

	api := []struct {
		MethodFunc func(pattern string, handlerFn http.HandlerFunc)
		Method     string
		Path       string
		HandleFunc func() func(w http.ResponseWriter, r *http.Request)
	}{
		{router.Post, "POST", "/evaluate", nh.Evaluate},
		{router.Post, "POST", "/validate", nh.Validate},
		{router.Get, "GET", "/errors", nh.Errors},
	}

	for _, endpoint := range api {
		endpoint.MethodFunc(endpoint.Path, endpoint.HandleFunc())
	}

	return router
}
