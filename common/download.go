package common

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func HttpDownload(rootURL, remotePath, localPath, filename string) error {
	_, err := HttpDownloadRename(rootURL, remotePath, localPath, filename, "")

	return err
}

func HttpDownloadRename(rootURL, remotePath, localPath, filename, rename string) (lastSavedArtPath string, error error) {
	logger := GetLoggerInstance()

	logger.Debug("Downloading file...",
		zap.String("remotePath", remotePath),
		zap.String("localPath", localPath),
		zap.String("filename", filename),
		zap.String("rename", rename))

	sourceURL, err := url.JoinPath(rootURL, remotePath, filename)
	if err != nil {
		return "", fmt.Errorf("unable to build download url: %w", err)
	}

	resp, err := http.Get(sourceURL)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	err = os.MkdirAll(localPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	fn := filename

	if rename != "" {
		// Used by the thumbnail downloader when a filename doesn't have the matching tags
		imageExt := filepath.Ext(filename)
		fn = strings.TrimSuffix(rename, filepath.Ext(rename))
		fn = fn + imageExt
	}

	f, err := os.Create(filepath.Join(localPath, fn))
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return filepath.Join(localPath, fn), nil
}
