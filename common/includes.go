package common

import (
	_ "embed"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

//go:embed resources/data/systems-mapping.json
var systemMapping []byte

func InitIncludes() {
	logger := GetLoggerInstance()
	cwd, _ := os.Getwd()

	dataPath := filepath.Join(cwd, "data")

	dataExists := false

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		err := os.MkdirAll(dataPath, 0755)
		if err != nil {
			logger.Fatal("Failed to create data directory", zap.Error(err))
		}
	} else {
		dataExists = true
	}

	if !dataExists {
		saveFile(systemMapping, filepath.Join(dataPath, "systems-mapping.json"))
	}
}

func saveFile(data []byte, path string) {
	logger := GetLoggerInstance()

	file, err := os.Create(path)
	if err != nil {
		logger.Fatal("Failed to open / create / truncate", zap.String("file_path", path), zap.Error(err))
	}

	_, err = file.Write(data)
	if err != nil {
		logger.Fatal("Failed to write", zap.String("file_path", path), zap.Error(err))
	}

}
