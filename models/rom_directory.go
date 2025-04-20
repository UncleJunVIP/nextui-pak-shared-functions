package models

type RomDirectory struct {
	DisplayName string
	Tag         string
	Path        string
}

func (r RomDirectory) Value() interface{} {
	return r
}

type RomDirectories []RomDirectory

func (r RomDirectories) Values() []string {
	var list []string
	for _, romDirectory := range r {
		list = append(list, romDirectory.DisplayName)
	}
	return list
}
