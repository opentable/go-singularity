package dtos

import "io"

type EmbeddedArtifact struct {
	Content                    StringList
	Filename                   string
	Md5sum                     string
	Name                       string
	TargetFolderRelativeToTask string
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

func (list EmbeddedArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list EmbeddedArtifactList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list EmbeddedArtifactList) FormatJSON() string {
	return FormatJSON(list)
}
