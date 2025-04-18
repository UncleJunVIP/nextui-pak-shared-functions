package common

import (
	_ "embed"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

//go:embed resources/data/systems-mapping.json
var systemMapping []byte

//go:embed resources/bin/tg5040/minui-keyboard
var keyboard []byte

//go:embed resources/bin/tg5040/minui-list
var list []byte

//go:embed resources/bin/tg5040/minui-presenter
var presenter []byte

func InitIncludes() {
	logger := GetLoggerInstance()
	cwd, _ := os.Getwd()

	binPath := filepath.Join(cwd, "bin/tg5040")
	dataPath := filepath.Join(cwd, "data")

	binExists := false
	dataExists := false

	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		err := os.MkdirAll(binPath, 0755)
		if err != nil {
			logger.Fatal("Failed to create bin directory", zap.Error(err))
		}
	} else {
		binExists = true
	}

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		err := os.MkdirAll(dataPath, 0755)
		if err != nil {
			logger.Fatal("Failed to create data directory", zap.Error(err))
		}
	} else {
		dataExists = true
	}

	if binExists && dataExists {
		return
	}

	if !binExists {
		saveFile(keyboard, filepath.Join(binPath, "minui-keyboard"))
		saveFile(list, filepath.Join(binPath, "minui-list"))
		saveFile(presenter, filepath.Join(binPath, "minui-presenter"))
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
