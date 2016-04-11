package dtos

import "io"

type SingularitySandbox struct {
	CurrentDirectory string
	Files            SingularitySandboxFileList
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

type SingularitySandboxList []*SingularitySandbox

func (list *SingularitySandboxList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularitySandboxList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularitySandboxList) FormatJSON() string {
	return FormatJSON(list)
}
