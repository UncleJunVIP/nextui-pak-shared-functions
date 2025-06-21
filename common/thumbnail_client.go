package common

import (
	"encoding/json"
	"fmt"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/models"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"qlova.tech/sum"
	"strings"
	"time"
)

const ThumbnailServerRoot = "https://thumbnails.libretro.com"

var InMemoryCache map[string]map[string]models.Items

func init() {
	if InMemoryCache == nil {
		InMemoryCache = make(map[string]map[string]models.Items)
	}
}

type ThumbnailClient struct {
	HttpTableClient
	SystemMapping   map[string]string
	ArtDownloadType sum.Int[models.ArtDownloadType]
}

func NewThumbnailClient(artDownloadType sum.Int[models.ArtDownloadType]) *ThumbnailClient {
	client := &ThumbnailClient{
		HttpTableClient: HttpTableClient{
			RootURL:  ThumbnailServerRoot,
			HostType: models.HostTypes.APACHE,
			TableColumns: models.TableColumns{
				FilenameHeader: "Name",
				FileSizeHeader: "",
				DateHeader:     "",
			},
		},
		ArtDownloadType: artDownloadType,
	}

	cwd, _ := os.Getwd()

	jsonPath := filepath.Join(cwd, "data", "systems-mapping.json")
	file, err := os.Open(jsonPath)
	if err == nil {
		defer file.Close()
		_ = json.NewDecoder(file).Decode(&client.SystemMapping)
	}

	return client
}

func (c *ThumbnailClient) BuildThumbnailSection(tag string) models.Section {
	systemName := c.SystemMapping[tag]
	subdirectory := path.Join(systemName, models.ArtDownloadTypeMapping[c.ArtDownloadType])

	return models.Section{
		Name:             tag,
		HostSubdirectory: subdirectory,
		LocalDirectory:   "",
	}
}

func (c *ThumbnailClient) Close() error {
	return nil
}

func (c *ThumbnailClient) ListDirectory(subdirectory string) (models.Items, error) {
	if cachedType, ok := InMemoryCache[c.ArtDownloadType.String()]; ok {
		if artList, ok := cachedType[subdirectory]; ok {
			return artList, nil
		}
	}

	artList, err := c.HttpTableClient.ListDirectory(subdirectory)

	if err != nil {
		return nil, fmt.Errorf("unable to list thumbnail directory: %w", err)
	}

	if _, ok := InMemoryCache[c.ArtDownloadType.String()]; !ok {
		InMemoryCache[c.ArtDownloadType.String()] = make(map[string]models.Items)
		InMemoryCache[c.ArtDownloadType.String()][subdirectory] = artList
	}

	return artList, nil
}

func (c *ThumbnailClient) BuildDownloadHeaders() map[string]string {
	headers := make(map[string]string)
	return headers
}

func (c *ThumbnailClient) DownloadArt(remotePath, localPath, filename, rename string) (savedPath string, error error) {
	logger := GetLoggerInstance()

	logger.Debug("Downloading file...",
		zap.String("remotePath", remotePath),
		zap.String("localPath", localPath),
		zap.String("filename", filename),
		zap.String("rename", rename))

	sourceURL, err := url.JoinPath(c.RootURL, remotePath, filename)
	if err != nil {
		return "", fmt.Errorf("unable to build download url: %w", err)
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := httpClient.Get(sourceURL)
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
		imageExt := filepath.Ext(filename)
		fn = strings.ReplaceAll(rename, filepath.Ext(rename), "")
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
