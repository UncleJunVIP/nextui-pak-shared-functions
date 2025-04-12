package common

import (
	"github.com/UncleJunVIP/nextui-pak-shared-functions/models"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const RomDirectory = "/mnt/SDCARD/Roms"

var TagRegex = regexp.MustCompile(`\((.*?)\)`)

func FetchRomDirectories() ([]models.RomDirectory, error) {
	entries, err := os.ReadDir(romDirectory)
	if err != nil {
		return nil, err
	}

	var dirs []models.RomDirectory

	for _, entry := range entries {
		if entry.IsDir() {
			tag := TagRegex.FindStringSubmatch(entry.Name())
			if tag == nil {
				continue
			}

			path := filepath.Join(romDirectory, entry.Name())
			tagless := strings.TrimSuffix(entry.Name(), tag[0])

			// For people that order their ROM directories
			if strings.Contains(tagless, ") ") {
				tagless = strings.Split(tagless, ") ")[1]
			}

			dirs = append(dirs, models.RomDirectory{
				DisplayName: strings.TrimSpace(tagless),
				Tag:         strings.TrimSpace(tag[1]),
				Path:        path,
			})
		}
	}

	return dirs, nil
}

func FetchRomDirectoriesByTag() (map[string]string, error) {
	dirs := make(map[string]string)

	entries, err := os.ReadDir(romDirectory)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			tag := TagRegex.FindStringSubmatch(entry.Name())
			if tag == nil {
				continue
			}

			path := filepath.Join(romDirectory, entry.Name())
			dirs[tag[1]] = path

		}
	}

	return dirs, nil
}
