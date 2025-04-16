package models

import "go.uber.org/zap/zapcore"

type Item struct {
	DisplayName string `json:"name"`

	Filename           string `json:"filename"`
	Path               string `json:"path"`
	IsDirectory        bool   `json:"is_directory"`
	DirectoryFileCount int    `json:"-"`
	FileSize           string `json:"file_size"`
	LastModified       string `json:"last_modified"`

	Tag string `json:"tag"`

	RomID  string `json:"-"` // For RomM Support
	ArtURL string `json:"-"` // For RomM Support
}

func (i Item) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("name", i.DisplayName)
	encoder.AddString("filename", i.Filename)
	encoder.AddString("path", i.Path)
	encoder.AddBool("is_directory", i.IsDirectory)
	encoder.AddString("file_size", i.FileSize)
	encoder.AddString("last_modified", i.LastModified)
	encoder.AddString("tag", i.Tag)
	encoder.AddString("rom_id", i.RomID)
	encoder.AddString("art_url", i.ArtURL)

	return nil
}

func (i Item) GetFilename() string {
	return i.Filename
}

type Items []Item

func (items Items) Values() []string {
	var list []string
	for _, item := range items {
		list = append(list, item.DisplayName)
	}
	return list
}

func (items Items) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, item := range items {
		_ = enc.AppendObject(item)
	}

	return nil
}
