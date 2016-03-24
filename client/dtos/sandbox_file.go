package dtos

import "io"

type SandboxFile struct {
	Mode string
	Mtime int64
	Name string
	Size int64
}

func (self *SandboxFile) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SandboxFile) FormatText() string {
	return FormatText(self)
}

func (self *SandboxFile) FormatJSON() string {
	return FormatJSON(self)
}
