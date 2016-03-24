package dtos

import "io"

type Sandbox struct {
	CurrentDirectory string
	Files []
	FullPathToRoot string
	SlaveHostname string
}

func (self *Sandbox) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Sandbox) FormatText() string {
	return FormatText(self)
}

func (self *Sandbox) FormatJSON() string {
	return FormatJSON(self)
}
