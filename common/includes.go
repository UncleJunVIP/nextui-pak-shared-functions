package common

import (
	"embed"
	"go.uber.org/zap"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed resources
var resources embed.FS

func InitIncludes() {
	logger := GetLoggerInstance()
	cwd, _ := os.Getwd()

	if err := os.MkdirAll(filepath.Join(cwd, "bin"), 0755); err != nil {
		logger.Fatal("Failed to create bin directory", zap.Error(err))
	}

	if err := os.MkdirAll(filepath.Join(cwd, "bin"), 0755); err != nil {
		logger.Fatal("Failed to create bin directory", zap.Error(err))
	}

	binaries, err := resources.ReadDir("resources/bin/tg5040")
	if err != nil {
		logger.Fatal("Failed to read bin directory", zap.Error(err))
	}

	dataFiles, err := resources.ReadDir("resources/data")
	if err != nil {
		logger.Fatal("Failed to read bin directory", zap.Error(err))
	}

	for _, file := range binaries {
		saveFile(file, resources, "resources/bin/tg5040/", "bin/tg5040")
	}

	for _, file := range dataFiles {
		saveFile(file, resources, "resources/data/", "data")
	}
}

func saveFile(file fs.DirEntry, resources embed.FS, resourceDir string, outDir string) {
	logger := GetLoggerInstance()
	cwd, _ := os.Getwd()

	bytes, err := resources.ReadFile(filepath.Join(resourceDir, file.Name()))
	if err != nil {
		logger.Fatal("Failed to read file", zap.String("resource_path", filepath.Join(resourceDir, file.Name())), zap.Error(err))
	}

	filePath := filepath.Join(cwd, outDir, file.Name())
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		logger.Fatal("Failed to write file", zap.String("file_path", filePath), zap.Error(err))
	}
}
