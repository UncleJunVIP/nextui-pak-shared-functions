package models

import "go.uber.org/zap/zapcore"

type Section struct {
	Name string `yaml:"section_name"`

	SystemTag      string `yaml:"system_tag"`
	LocalDirectory string `yaml:"local_directory"`

	HostSubdirectory string `yaml:"host_subdirectory"`

	RomMPlatformID string `yaml:"romm_platform_id"`
}

type Sections []Section

func (s Sections) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	for _, section := range s {
		_ = encoder.AppendObject(section)
	}

	return nil
}

func (s Section) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("name", s.Name)
	encoder.AddString("system_tag", s.SystemTag)
	encoder.AddString("local_directory", s.LocalDirectory)
	encoder.AddString("host_subdirectory", s.HostSubdirectory)
	encoder.AddString("romm_platform_id", s.RomMPlatformID)

	return nil
}
