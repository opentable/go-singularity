package dtos

import "io"

type S3ArtifactSignature struct {
	ArtifactFilename           string `json:"artifactFilename"`
	Filename                   string `json:"filename"`
	Filesize                   int64  `json:"filesize"`
	Md5sum                     string `json:"md5sum"`
	Name                       string `json:"name"`
	S3Bucket                   string `json:"s3Bucket"`
	S3ObjectKey                string `json:"s3ObjectKey"`
	TargetFolderRelativeToTask string `json:"targetFolderRelativeToTask"`
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

type S3ArtifactSignatureList []*S3ArtifactSignature

func (list *S3ArtifactSignatureList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *S3ArtifactSignatureList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *S3ArtifactSignatureList) FormatJSON() string {
	return FormatJSON(list)
}
