package filebrowser

import (
	"fmt"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/common"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/models"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/ui"
	"go.uber.org/zap"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FileBrowser struct {
	logger           *zap.Logger
	WorkingDirectory string
	Items            models.Items
	HumanReadableLS  map[string]models.Item
}

func NewFileBrowser(logger *zap.Logger) *FileBrowser {
	return &FileBrowser{
		logger: logger,
	}
}

func (c *FileBrowser) CWD(newDirectory string) error {
	files, err := os.ReadDir(newDirectory)
	if err != nil {
		return fmt.Errorf("failed to read directory %w", err)
	}

	c.WorkingDirectory = newDirectory
	updatedHumanReadable := make(map[string]models.Item)

	var items []models.Item
	for _, file := range files {
		displayName, tag := itemNameCleaner(file.Name(), false)

		directoryFileCount := -1
		if file.IsDir() {
			dir, err := os.ReadDir(path.Join(c.WorkingDirectory, file.Name()))
			if err != nil {
				c.logger.Error("Failed to read directory", zap.String("path", path.Join(c.WorkingDirectory, file.Name())), zap.Error(err))
			}

			directoryFileCount = len(dir)
		}

		item := models.Item{
			DisplayName:        displayName,
			Tag:                tag,
			Filename:           file.Name(),
			Path:               path.Join(c.WorkingDirectory, file.Name()),
			IsDirectory:        file.IsDir(),
			DirectoryFileCount: directoryFileCount,
		}

		if !file.IsDir() || directoryFileCount > 0 { // Hide them empty directories the lunatics leave behind...
			items = append(items, item)
			updatedHumanReadable[displayName] = item
		}
	}

	c.Items = items
	c.HumanReadableLS = updatedHumanReadable

	return nil
}

func (c *FileBrowser) DisplayCurrentDirectory(title string) (models.Item, error) {
	res, err := ui.DisplayList(c.Items, title, "")

	if err != nil {
		return models.Item{}, err
	}

	return c.HumanReadableLS[res.Value], nil
}

func itemNameCleaner(filename string, stripTag bool) (string, string) {
	cleaned := filepath.Clean(filename)

	// Clean up the tags
	tag := common.TagRegex.FindStringSubmatch(cleaned)

	foundTag := tag[1]

	if len(tag) > 0 && stripTag {
		cleaned = strings.TrimSuffix(filename, tag[0])
	}

	// For people that order their ROM directories
	if strings.Contains(cleaned, ") ") {
		cleaned = strings.Split(cleaned, ") ")[1]
	}

	// Lose the extension
	cleaned = strings.ReplaceAll(cleaned, path.Ext(cleaned), "")

	cleaned = strings.TrimSpace(cleaned)

	return cleaned, foundTag
}
