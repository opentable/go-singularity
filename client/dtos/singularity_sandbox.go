package dtos

import "io"

type SingularitySandbox struct {
	CurrentDirectory string
	Files            *SingularitySandboxFile
	FullPathToRoot   string
	SlaveHostname    string
}

func (self *SingularitySandbox) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularitySandbox) FormatText() string {
	return FormatText(self)
}

func (self *SingularitySandbox) FormatJSON() string {
	return FormatJSON(self)
}
