package dtos

import "io"

type TaskShellCommandUpdate struct {
	Message string
	OutputFilename string
	ShellRequestId TaskShellCommandRequestId
	Timestamp int64
//	UpdateType UpdateType
}

func (self *TaskShellCommandUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskShellCommandUpdate) FormatText() string {
	return FormatText(self)
}

func (self *TaskShellCommandUpdate) FormatJSON() string {
	return FormatJSON(self)
}
