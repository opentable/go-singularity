package dtos

import "io"

type EmbeddedArtifact struct {
	Content                    StringList `json:"content"`
	Filename                   string     `json:"filename"`
	Md5sum                     string     `json:"md5sum"`
	Name                       string     `json:"name"`
	TargetFolderRelativeToTask string     `json:"targetFolderRelativeToTask"`
}

func (self *EmbeddedArtifact) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *EmbeddedArtifact) FormatText() string {
	return FormatText(self)
}

func (self *EmbeddedArtifact) FormatJSON() string {
	return FormatJSON(self)
}

type EmbeddedArtifactList []*EmbeddedArtifact

func (list *EmbeddedArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *EmbeddedArtifactList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *EmbeddedArtifactList) FormatJSON() string {
	return FormatJSON(list)
}
