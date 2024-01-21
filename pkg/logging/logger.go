package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/vanya/backend/internal/config"
)

func NewLogger(env string) *logrus.Logger {
	log := logrus.New()

	switch env {
	case config.EnvLocal:
		log.SetFormatter(&logrus.TextFormatter{})
		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(os.Stdout)
	case config.EnvProd:
		log.SetFormatter(&logrus.JSONFormatter{})
		log.SetLevel(logrus.InfoLevel)
		log.SetOutput(os.Stdout)
	}

	return log
}
