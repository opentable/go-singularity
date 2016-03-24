package dtos

import "io"

type EmbeddedArtifact struct {
	Content []string
	Filename string
	Md5sum string
	Name string
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
