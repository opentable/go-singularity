package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type S3Artifact struct {
	present map[string]bool

	TargetFolderRelativeToTask string `json:"targetFolderRelativeToTask,omitempty"`

	Name string `json:"name,omitempty"`

	S3Bucket string `json:"s3Bucket,omitempty"`

	S3ObjectKey string `json:"s3ObjectKey,omitempty"`

	Filesize int64 `json:"filesize"`

	IsArtifactList bool `json:"isArtifactList"`

	Filename string `json:"filename,omitempty"`

	Md5sum string `json:"md5sum,omitempty"`
}

func (self *S3Artifact) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *S3Artifact) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*S3Artifact); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A S3Artifact cannot copy the values from %#v", other)
}

func (self *S3Artifact) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *S3Artifact) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *S3Artifact) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *S3Artifact) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *S3Artifact) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on S3Artifact", name)

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		v, ok := value.(string)
		if ok {
			self.TargetFolderRelativeToTask = v
			self.present["targetFolderRelativeToTask"] = true
			return nil
		} else {
			return fmt.Errorf("Field targetFolderRelativeToTask/TargetFolderRelativeToTask: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = v
			self.present["name"] = true
			return nil
		} else {
			return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "s3Bucket", "S3Bucket":
		v, ok := value.(string)
		if ok {
			self.S3Bucket = v
			self.present["s3Bucket"] = true
			return nil
		} else {
			return fmt.Errorf("Field s3Bucket/S3Bucket: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "s3ObjectKey", "S3ObjectKey":
		v, ok := value.(string)
		if ok {
			self.S3ObjectKey = v
			self.present["s3ObjectKey"] = true
			return nil
		} else {
			return fmt.Errorf("Field s3ObjectKey/S3ObjectKey: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "filesize", "Filesize":
		v, ok := value.(int64)
		if ok {
			self.Filesize = v
			self.present["filesize"] = true
			return nil
		} else {
			return fmt.Errorf("Field filesize/Filesize: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "isArtifactList", "IsArtifactList":
		v, ok := value.(bool)
		if ok {
			self.IsArtifactList = v
			self.present["isArtifactList"] = true
			return nil
		} else {
			return fmt.Errorf("Field isArtifactList/IsArtifactList: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "filename", "Filename":
		v, ok := value.(string)
		if ok {
			self.Filename = v
			self.present["filename"] = true
			return nil
		} else {
			return fmt.Errorf("Field filename/Filename: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "md5sum", "Md5sum":
		v, ok := value.(string)
		if ok {
			self.Md5sum = v
			self.present["md5sum"] = true
			return nil
		} else {
			return fmt.Errorf("Field md5sum/Md5sum: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *S3Artifact) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on S3Artifact", name)

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		if self.present != nil {
			if _, ok := self.present["targetFolderRelativeToTask"]; ok {
				return self.TargetFolderRelativeToTask, nil
			}
		}
		return nil, fmt.Errorf("Field TargetFolderRelativeToTask no set on TargetFolderRelativeToTask %+v", self)

	case "name", "Name":
		if self.present != nil {
			if _, ok := self.present["name"]; ok {
				return self.Name, nil
			}
		}
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "s3Bucket", "S3Bucket":
		if self.present != nil {
			if _, ok := self.present["s3Bucket"]; ok {
				return self.S3Bucket, nil
			}
		}
		return nil, fmt.Errorf("Field S3Bucket no set on S3Bucket %+v", self)

	case "s3ObjectKey", "S3ObjectKey":
		if self.present != nil {
			if _, ok := self.present["s3ObjectKey"]; ok {
				return self.S3ObjectKey, nil
			}
		}
		return nil, fmt.Errorf("Field S3ObjectKey no set on S3ObjectKey %+v", self)

	case "filesize", "Filesize":
		if self.present != nil {
			if _, ok := self.present["filesize"]; ok {
				return self.Filesize, nil
			}
		}
		return nil, fmt.Errorf("Field Filesize no set on Filesize %+v", self)

	case "isArtifactList", "IsArtifactList":
		if self.present != nil {
			if _, ok := self.present["isArtifactList"]; ok {
				return self.IsArtifactList, nil
			}
		}
		return nil, fmt.Errorf("Field IsArtifactList no set on IsArtifactList %+v", self)

	case "filename", "Filename":
		if self.present != nil {
			if _, ok := self.present["filename"]; ok {
				return self.Filename, nil
			}
		}
		return nil, fmt.Errorf("Field Filename no set on Filename %+v", self)

	case "md5sum", "Md5sum":
		if self.present != nil {
			if _, ok := self.present["md5sum"]; ok {
				return self.Md5sum, nil
			}
		}
		return nil, fmt.Errorf("Field Md5sum no set on Md5sum %+v", self)

	}
}

func (self *S3Artifact) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on S3Artifact", name)

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		self.present["targetFolderRelativeToTask"] = false

	case "name", "Name":
		self.present["name"] = false

	case "s3Bucket", "S3Bucket":
		self.present["s3Bucket"] = false

	case "s3ObjectKey", "S3ObjectKey":
		self.present["s3ObjectKey"] = false

	case "filesize", "Filesize":
		self.present["filesize"] = false

	case "isArtifactList", "IsArtifactList":
		self.present["isArtifactList"] = false

	case "filename", "Filename":
		self.present["filename"] = false

	case "md5sum", "Md5sum":
		self.present["md5sum"] = false

	}

	return nil
}

func (self *S3Artifact) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type S3ArtifactList []*S3Artifact

func (self *S3ArtifactList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*S3ArtifactList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A S3ArtifactList cannot copy the values from %#v", other)
}

func (list *S3ArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
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
	return swaggering.FormatJSON(list)
}
