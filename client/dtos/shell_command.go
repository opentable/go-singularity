package dtos

import "io"

type ShellCommand struct {
	LogfileName string
	Name string
	Options []string
	User string
}

func (self *ShellCommand) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ShellCommand) FormatText() string {
	return FormatText(self)
}

func (self *ShellCommand) FormatJSON() string {
	return FormatJSON(self)
}
