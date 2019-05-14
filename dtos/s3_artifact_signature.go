package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type S3ArtifactSignature struct {
	S3ObjectKey                *string `json:"s3ObjectKey,omitempty"`
	Filesize                   *int64  `json:"filesize,omitempty"`
	IsArtifactList             *bool   `json:"isArtifactList,omitempty"`
	Filename                   *string `json:"filename,omitempty"`
	ArtifactFilename           *string `json:"artifactFilename,omitempty"`
	S3Bucket                   *string `json:"s3Bucket,omitempty"`
	Md5sum                     *string `json:"md5sum,omitempty"`
	TargetFolderRelativeToTask *string `json:"targetFolderRelativeToTask,omitempty"`
	Name                       *string `json:"name,omitempty"`
}

func (self *S3ArtifactSignature) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *S3ArtifactSignature) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*S3ArtifactSignature); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A S3ArtifactSignature cannot copy the values from %#v", other)
}

func (self *S3ArtifactSignature) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *S3ArtifactSignature) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *S3ArtifactSignature) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on S3ArtifactSignature", name)

	case "s3ObjectKey", "S3ObjectKey":
		v, ok := value.(string)
		if ok {
			self.S3ObjectKey = &v
			return nil
		}
		return fmt.Errorf("Field s3ObjectKey/S3ObjectKey: value %v(%T) couldn't be cast to type string", value, value)

	case "filesize", "Filesize":
		v, ok := value.(int64)
		if ok {
			self.Filesize = &v
			return nil
		}
		return fmt.Errorf("Field filesize/Filesize: value %v(%T) couldn't be cast to type int64", value, value)

	case "isArtifactList", "IsArtifactList":
		v, ok := value.(bool)
		if ok {
			self.IsArtifactList = &v
			return nil
		}
		return fmt.Errorf("Field isArtifactList/IsArtifactList: value %v(%T) couldn't be cast to type bool", value, value)

	case "filename", "Filename":
		v, ok := value.(string)
		if ok {
			self.Filename = &v
			return nil
		}
		return fmt.Errorf("Field filename/Filename: value %v(%T) couldn't be cast to type string", value, value)

	case "artifactFilename", "ArtifactFilename":
		v, ok := value.(string)
		if ok {
			self.ArtifactFilename = &v
			return nil
		}
		return fmt.Errorf("Field artifactFilename/ArtifactFilename: value %v(%T) couldn't be cast to type string", value, value)

	case "s3Bucket", "S3Bucket":
		v, ok := value.(string)
		if ok {
			self.S3Bucket = &v
			return nil
		}
		return fmt.Errorf("Field s3Bucket/S3Bucket: value %v(%T) couldn't be cast to type string", value, value)

	case "md5sum", "Md5sum":
		v, ok := value.(string)
		if ok {
			self.Md5sum = &v
			return nil
		}
		return fmt.Errorf("Field md5sum/Md5sum: value %v(%T) couldn't be cast to type string", value, value)

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		v, ok := value.(string)
		if ok {
			self.TargetFolderRelativeToTask = &v
			return nil
		}
		return fmt.Errorf("Field targetFolderRelativeToTask/TargetFolderRelativeToTask: value %v(%T) couldn't be cast to type string", value, value)

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = &v
			return nil
		}
		return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *S3ArtifactSignature) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on S3ArtifactSignature", name)

	case "s3ObjectKey", "S3ObjectKey":
		return *self.S3ObjectKey, nil
		return nil, fmt.Errorf("Field S3ObjectKey no set on S3ObjectKey %+v", self)

	case "filesize", "Filesize":
		return *self.Filesize, nil
		return nil, fmt.Errorf("Field Filesize no set on Filesize %+v", self)

	case "isArtifactList", "IsArtifactList":
		return *self.IsArtifactList, nil
		return nil, fmt.Errorf("Field IsArtifactList no set on IsArtifactList %+v", self)

	case "filename", "Filename":
		return *self.Filename, nil
		return nil, fmt.Errorf("Field Filename no set on Filename %+v", self)

	case "artifactFilename", "ArtifactFilename":
		return *self.ArtifactFilename, nil
		return nil, fmt.Errorf("Field ArtifactFilename no set on ArtifactFilename %+v", self)

	case "s3Bucket", "S3Bucket":
		return *self.S3Bucket, nil
		return nil, fmt.Errorf("Field S3Bucket no set on S3Bucket %+v", self)

	case "md5sum", "Md5sum":
		return *self.Md5sum, nil
		return nil, fmt.Errorf("Field Md5sum no set on Md5sum %+v", self)

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		return *self.TargetFolderRelativeToTask, nil
		return nil, fmt.Errorf("Field TargetFolderRelativeToTask no set on TargetFolderRelativeToTask %+v", self)

	case "name", "Name":
		return *self.Name, nil
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	}
}

func (self *S3ArtifactSignature) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on S3ArtifactSignature", name)

	case "s3ObjectKey", "S3ObjectKey":
		self.S3ObjectKey = nil

	case "filesize", "Filesize":
		self.Filesize = nil

	case "isArtifactList", "IsArtifactList":
		self.IsArtifactList = nil

	case "filename", "Filename":
		self.Filename = nil

	case "artifactFilename", "ArtifactFilename":
		self.ArtifactFilename = nil

	case "s3Bucket", "S3Bucket":
		self.S3Bucket = nil

	case "md5sum", "Md5sum":
		self.Md5sum = nil

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		self.TargetFolderRelativeToTask = nil

	case "name", "Name":
		self.Name = nil

	}

	return nil
}

func (self *S3ArtifactSignature) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type S3ArtifactSignatureList []*S3ArtifactSignature

func (self *S3ArtifactSignatureList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*S3ArtifactSignatureList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A S3ArtifactSignatureList cannot copy the values from %#v", other)
}

func (list *S3ArtifactSignatureList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
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
	return swaggering.FormatJSON(list)
}
