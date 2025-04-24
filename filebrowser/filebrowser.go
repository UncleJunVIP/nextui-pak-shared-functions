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
		displayName, tag := ItemNameCleaner(file.Name(), false)

		directoryFileCount := -1
		isMultiDisc := false
		if file.IsDir() {
			dir, err := os.ReadDir(path.Join(c.WorkingDirectory, file.Name()))
			if err != nil {
				c.logger.Error("Failed to read directory", zap.String("path", path.Join(c.WorkingDirectory, file.Name())), zap.Error(err))
			}

			for _, f := range dir {
				if strings.Contains(f.Name(), "Disc") ||
					strings.Contains(f.Name(), "Disk") ||
					strings.Contains(f.Name(), "CD") ||
					strings.Contains(f.Name(), ".bin") ||
					strings.Contains(f.Name(), ".cue") {
					isMultiDisc = true
					break
				}
			}

			directoryFileCount = len(dir)
		}

		item := models.Item{
			DisplayName:          displayName,
			Tag:                  tag,
			Filename:             file.Name(),
			Path:                 path.Join(c.WorkingDirectory, file.Name()),
			IsDirectory:          !isMultiDisc && file.IsDir(),
			IsMultiDiscDirectory: isMultiDisc,
			DirectoryFileCount:   directoryFileCount,
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

	return c.HumanReadableLS[res.SelectedValue], nil
}

func ItemNameCleaner(filename string, stripTag bool) (string, string) {
	cleaned := filepath.Clean(filename)

	tags := common.TagRegex.FindAllStringSubmatch(cleaned, -1)

	var foundTags []string
	foundTag := ""

	if len(tags) > 0 {
		for _, tagPair := range tags {
			foundTags = append(foundTags, tagPair[0])
		}

		foundTag = strings.Join(foundTags, " ")

		if stripTag {
			cleaned = strings.ReplaceAll(filename, foundTag, "")
		}

	}

	orderedFolderRegex := common.OrderedFolderRegex.FindStringSubmatch(cleaned)

	if len(orderedFolderRegex) > 0 {
		cleaned = strings.ReplaceAll(cleaned, orderedFolderRegex[0], "")
	}

	// Lose the extension
	cleaned = strings.ReplaceAll(cleaned, path.Ext(cleaned), "")

	cleaned = strings.TrimSpace(cleaned)

	return cleaned, foundTag
}
