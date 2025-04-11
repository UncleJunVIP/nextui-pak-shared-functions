package models

import "go.uber.org/zap/zapcore"

type Item struct {
	Filename string `json:"filename"`
}

func (i Item) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("filename", i.Filename)

	return nil
}

type Items []Item

func (i Items) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, item := range i {
		_ = enc.AppendObject(item)
	}

	return nil
}
