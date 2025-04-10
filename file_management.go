package nextuisharedfunctions

import (
	"os"
	"path/filepath"
	"regexp"
)

const romDirectory = "/mnt/SDCARD/Roms"

var TagRegex = regexp.MustCompile(`\((.*?)\)`)

func FetchRomDirectories() (map[string]string, error) {
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
