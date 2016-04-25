package dtos

import "io"

type S3Artifact struct {
	Filename                   string `json:"filename"`
	Filesize                   int64  `json:"filesize"`
	Md5sum                     string `json:"md5sum"`
	Name                       string `json:"name"`
	S3Bucket                   string `json:"s3Bucket"`
	S3ObjectKey                string `json:"s3ObjectKey"`
	TargetFolderRelativeToTask string `json:"targetFolderRelativeToTask"`
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

type S3ArtifactList []*S3Artifact

func (list *S3ArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *S3ArtifactList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *S3ArtifactList) FormatJSON() string {
	return FormatJSON(list)
}
