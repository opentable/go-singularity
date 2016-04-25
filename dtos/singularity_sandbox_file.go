package dtos

import "io"

type SingularitySandboxFile struct {
	Mode  string `json:"mode"`
	Mtime int64  `json:"mtime"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
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

type SingularitySandboxFileList []*SingularitySandboxFile

func (list *SingularitySandboxFileList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularitySandboxFileList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularitySandboxFileList) FormatJSON() string {
	return FormatJSON(list)
}
