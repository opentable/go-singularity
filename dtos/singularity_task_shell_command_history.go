package dtos

import "io"

type SingularityTaskShellCommandHistory struct {
	ShellRequest *SingularityTaskShellCommandRequest
	ShellUpdates *SingularityTaskShellCommandUpdate
}

func (self *SingularityTaskShellCommandHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskShellCommandHistory) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskShellCommandHistory) FormatJSON() string {
	return FormatJSON(self)
}
