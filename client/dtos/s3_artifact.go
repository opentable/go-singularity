package dtos

import "io"

type S3Artifact struct {
	Filename                   string
	Filesize                   int64
	Md5sum                     string
	Name                       string
	S3Bucket                   string
	S3ObjectKey                string
	TargetFolderRelativeToTask string
}

func (self *S3Artifact) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *S3Artifact) FormatText() string {
	return FormatText(self)
}

func (self *S3Artifact) FormatJSON() string {
	return FormatJSON(self)
}
