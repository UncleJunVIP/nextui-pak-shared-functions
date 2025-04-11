package models

import "go.uber.org/zap/zapcore"

type TableColumns struct {
	FilenameHeader string `yaml:"filename_header"`
	FileSizeHeader string `yaml:"file_size_header"`
	DateHeader     string `yaml:"date_header"`
}

func (c TableColumns) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("filename_header", c.FilenameHeader)
	enc.AddString("file_size_header", c.FileSizeHeader)
	enc.AddString("date_header", c.DateHeader)

	return nil
}
