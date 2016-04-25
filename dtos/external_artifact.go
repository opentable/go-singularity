package dtos

import "io"

type ExternalArtifact struct {
	Filename                   string `json:"filename"`
	Filesize                   int64  `json:"filesize"`
	Md5sum                     string `json:"md5sum"`
	Name                       string `json:"name"`
	TargetFolderRelativeToTask string `json:"targetFolderRelativeToTask"`
	Url                        string `json:"url"`
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

func (list *ExternalArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ExternalArtifactList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ExternalArtifactList) FormatJSON() string {
	return FormatJSON(list)
}
