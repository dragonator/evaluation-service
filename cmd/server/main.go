package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dragonator/evaluation-service/module/evaluation"
	"github.com/dragonator/evaluation-service/pkg/config"
	"github.com/dragonator/evaluation-service/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(cfg.LoggerLevel)

	serverModule, err := evaluation.NewServerModule(cfg, logger)
	if err != nil {
		panic(err)
	}

	serverModule.EvaluationService.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop

	log.Printf("Signal caught (%s), stopping...", sig.String())
	serverModule.EvaluationService.Stop()
	log.Print("Service stopped.")
}
