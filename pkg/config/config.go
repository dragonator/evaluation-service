// Package config contains configuration structs.
package config

import (
	"errors"
	"fmt"
	"os"
)

var _errUndefinedEnvVar = errors.New("undefined environment variable")

// Config holds the service config.
type Config struct {
	ServerPort  string
	LoggerLevel string
}

// New is a constructor function for Config.
func New() (*Config, error) {
	serverPort, defined := os.LookupEnv("SERVER_PORT")
	if !defined {
		return nil, fmt.Errorf("%w: SERVER_PORT", _errUndefinedEnvVar)
	}

	loggerLevel, defined := os.LookupEnv("LOGGER_LEVEL")
	if !defined {
		return nil, fmt.Errorf("%w: LOGGER_LEVEL", _errUndefinedEnvVar)
	}

	return &Config{
		ServerPort:  serverPort,
		LoggerLevel: loggerLevel,
	}, nil
}
