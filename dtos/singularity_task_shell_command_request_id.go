package dtos

import "io"

type SingularityTaskShellCommandRequestId struct {
	Id        string
	Name      string
	TaskId    *SingularityTaskId
	Timestamp int64
}

func (self *SingularityTaskShellCommandRequestId) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskShellCommandRequestId) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskShellCommandRequestId) FormatJSON() string {
	return FormatJSON(self)
}
