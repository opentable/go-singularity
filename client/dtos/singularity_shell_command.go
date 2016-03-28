package dtos

import "io"

type SingularityShellCommand struct {
	LogfileName string
	Name        string
	Options     string
	User        string
}

func (self *SingularityShellCommand) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityShellCommand) FormatText() string {
	return FormatText(self)
}

func (self *SingularityShellCommand) FormatJSON() string {
	return FormatJSON(self)
}
