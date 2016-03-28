package dtos

import "io"

type S3ArtifactSignature struct {
	ArtifactFilename           string
	Filename                   string
	Filesize                   int64
	Md5sum                     string
	Name                       string
	S3Bucket                   string
	S3ObjectKey                string
	TargetFolderRelativeToTask string
}

func (self *S3ArtifactSignature) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *S3ArtifactSignature) FormatText() string {
	return FormatText(self)
}

func (self *S3ArtifactSignature) FormatJSON() string {
	return FormatJSON(self)
}
