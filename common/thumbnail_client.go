package common

import (
	"encoding/json"
	"fmt"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/models"
	"os"
	"path"
	"path/filepath"
	"qlova.tech/sum"
)

const ThumbnailServerRoot = "https://thumbnails.libretro.com"

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

func (c *ThumbnailClient) ListDirectory(section models.Section) ([]models.Item, error) {
	artList, err := c.HttpTableClient.ListDirectory(section.HostSubdirectory)

	if err != nil {
		return nil, fmt.Errorf("unable to list thumbnail directory: %w", err)
	}

	return artList, nil
}

func (c *ThumbnailClient) DownloadFile(remotePath, localPath, filename string) (lastSavedPath string, error error) {
	return HttpDownload(c.RootURL, remotePath, localPath, filename)
}

func (c *ThumbnailClient) DownloadFileRename(remotePath, localPath, filename, rename string) (lastSavedPath string, error error) {
	return HttpDownloadRename(c.RootURL, remotePath, localPath, filename, rename)
}
