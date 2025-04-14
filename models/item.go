package models

import "go.uber.org/zap/zapcore"

type Item struct {
	DisplayName string `json:"name"`
	Tag         string `json:"tag"`
	Filename    string `json:"filename"`
	Path        string `json:"path"`
	IsDirectory bool   `json:"is_directory"`
}

func (i Item) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("name", i.DisplayName)
	encoder.AddString("tag", i.Tag)
	encoder.AddString("filename", i.Filename)
	encoder.AddString("path", i.Path)
	encoder.AddBool("is_directory", i.IsDirectory)

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
