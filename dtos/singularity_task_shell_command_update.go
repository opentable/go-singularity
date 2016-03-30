package dtos

import "io"

type SingularityTaskShellCommandUpdate struct {
	Message        string
	OutputFilename string
	ShellRequestId *SingularityTaskShellCommandRequestId
	Timestamp      int64
	//	UpdateType *UpdateType
}

func (self *SingularityTaskShellCommandUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskShellCommandUpdate) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskShellCommandUpdate) FormatJSON() string {
	return FormatJSON(self)
}
