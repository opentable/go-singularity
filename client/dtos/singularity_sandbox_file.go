package dtos

import "io"

type SingularitySandboxFile struct {
	Mode  string
	Mtime int64
	Name  string
	Size  int64
}

func (self *SingularitySandboxFile) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularitySandboxFile) FormatText() string {
	return FormatText(self)
}

func (self *SingularitySandboxFile) FormatJSON() string {
	return FormatJSON(self)
}
