package dtos

import "io"

type ExternalArtifact struct {
	Filename                   string
	Filesize                   int64
	Md5sum                     string
	Name                       string
	TargetFolderRelativeToTask string
	Url                        string
}

func (self *ExternalArtifact) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExternalArtifact) FormatText() string {
	return FormatText(self)
}

func (self *ExternalArtifact) FormatJSON() string {
	return FormatJSON(self)
}

type ExternalArtifactList []*ExternalArtifact

func (list ExternalArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list ExternalArtifactList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list ExternalArtifactList) FormatJSON() string {
	return FormatJSON(list)
}
