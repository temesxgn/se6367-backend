package logger

import (
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
	"github.com/temesxgn/se6367-backend/config"
	"github.com/temesxgn/se6367-backend/metrics"
)

// CreateLogger - creates a new log entry with the formatter, level and app name set
func CreateLogger(appName string) *logrus.Entry {
	log := logrus.New()
	childFormatter := logrus.JSONFormatter{}
	runtimeFormatter := &runtime.Formatter{ChildFormatter: &childFormatter}
	log.Formatter = runtimeFormatter

	// Set logger level and out based on mode
	mode := config.GetApplicationMode()
	switch mode {
	case config.ProdMode:
		log.SetLevel(logrus.InfoLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}

	entry := log.WithField(metrics.AppName, appName)
	if mode != "" {
		entry = entry.WithField(metrics.Mode, mode.String())
	}

	return entry
}
