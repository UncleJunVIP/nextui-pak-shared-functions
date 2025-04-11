package models

import "go.uber.org/zap/zapcore"

type Section struct {
	Name string `yaml:"section_name"`

	SystemTag      string `yaml:"system_tag"`
	LocalDirectory string `yaml:"local_directory"`

	HostSubdirectory string `yaml:"host_subdirectory"`
}

type Sections []Section

func (s Sections) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, section := range s {
		_ = enc.AppendObject(section)
	}

	return nil
}

func (s Section) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", s.Name)
	enc.AddString("system_tag", s.SystemTag)
	enc.AddString("local_directory", s.LocalDirectory)
	enc.AddString("host_subdirectory", s.HostSubdirectory)

	return nil
}
