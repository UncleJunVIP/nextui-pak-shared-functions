package common

import (
	"go.uber.org/zap"
	"os"
)

func DeleteFile(path string) bool {
	logger := GetLoggerInstance()

	err := os.RemoveAll(path)
	if err != nil {
		logger.Error("Issue removing file",
			zap.String("path", path),
			zap.Error(err))
		return false
	} else {
		logger.Debug("Removed file", zap.String("path", path))
		return true
	}
}
