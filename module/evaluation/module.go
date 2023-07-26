package evaluation

import (
	"fmt"

	"github.com/dragonator/evaluation-service/module/evaluation/internal/http/handler"
	"github.com/dragonator/evaluation-service/module/evaluation/internal/http/service"
	"github.com/dragonator/evaluation-service/module/evaluation/internal/operation/evaluationprocessing"
	"github.com/dragonator/evaluation-service/pkg/config"
	"github.com/dragonator/evaluation-service/pkg/logger"
)

const _evaluationTopic = "evaluation_service_evaluations"

// NotificationService provides methods for starting and stopping a evaluation service.
type EvaluationService interface {
	Start()
	Stop()
}

// ServerModule provides access to the functionality of evaluation server module.
type ServerModule struct {
	EvaluationService EvaluationService
}

// NewServerModule is a construction function for ServerModule.
func NewServerModule(config *config.Config, logger *logger.Logger) (*ServerModule, error) {
	evaluationProcessingOp := evaluationprocessing.NewOperation()
	evaluationHandler := handler.NewEvaluationHandler(evaluationProcessingOp)
	router := service.NewRouter(evaluationHandler)

	evaluationService, err := service.New(config, logger, router)
	if err != nil {
		return nil, fmt.Errorf("creating evaluation module: %w", err)
	}

	return &ServerModule{
		EvaluationService: evaluationService,
	}, nil
}
