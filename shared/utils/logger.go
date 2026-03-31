package utils

import "go.uber.org/zap"

func Logger() *zap.SugaredLogger {

	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	return logger
}
