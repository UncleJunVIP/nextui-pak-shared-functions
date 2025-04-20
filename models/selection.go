package models

type ListSelection struct {
	SelectedValue string
	Index         int
	ExitCode      int
}

func (s ListSelection) Value() interface{} {
	return s
}
