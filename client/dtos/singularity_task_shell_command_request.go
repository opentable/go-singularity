package dtos

import "io"

type SingularityTaskShellCommandRequest struct {
	ShellCommand *SingularityShellCommand
	TaskId       *SingularityTaskId
	Timestamp    int64
	User         string
}

func (self *SingularityTaskShellCommandRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskShellCommandRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskShellCommandRequest) FormatJSON() string {
	return FormatJSON(self)
}
