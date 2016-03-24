package dtos

import "io"

type TaskShellCommandHistory struct {
	ShellRequest TaskShellCommandRequest
	ShellUpdates []
}

func (self *TaskShellCommandHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskShellCommandHistory) FormatText() string {
	return FormatText(self)
}

func (self *TaskShellCommandHistory) FormatJSON() string {
	return FormatJSON(self)
}
