package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ExternalArtifact struct {
	Name                       *string `json:"name,omitempty"`
	Url                        *string `json:"url,omitempty"`
	Filesize                   *int64  `json:"filesize,omitempty"`
	IsArtifactList             *bool   `json:"isArtifactList,omitempty"`
	Filename                   *string `json:"filename,omitempty"`
	Md5sum                     *string `json:"md5sum,omitempty"`
	TargetFolderRelativeToTask *string `json:"targetFolderRelativeToTask,omitempty"`
}

func (self *ExternalArtifact) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *ExternalArtifact) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ExternalArtifact); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ExternalArtifact cannot copy the values from %#v", other)
}

func (self *ExternalArtifact) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ExternalArtifact) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ExternalArtifact) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExternalArtifact", name)

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = &v
			return nil
		}
		return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)

	case "url", "Url":
		v, ok := value.(string)
		if ok {
			self.Url = &v
			return nil
		}
		return fmt.Errorf("Field url/Url: value %v(%T) couldn't be cast to type string", value, value)

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

	}
}

func (self *ExternalArtifact) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ExternalArtifact", name)

	case "name", "Name":
		return *self.Name, nil
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "url", "Url":
		return *self.Url, nil
		return nil, fmt.Errorf("Field Url no set on Url %+v", self)

	case "filesize", "Filesize":
		return *self.Filesize, nil
		return nil, fmt.Errorf("Field Filesize no set on Filesize %+v", self)

	case "isArtifactList", "IsArtifactList":
		return *self.IsArtifactList, nil
		return nil, fmt.Errorf("Field IsArtifactList no set on IsArtifactList %+v", self)

	case "filename", "Filename":
		return *self.Filename, nil
		return nil, fmt.Errorf("Field Filename no set on Filename %+v", self)

	case "md5sum", "Md5sum":
		return *self.Md5sum, nil
		return nil, fmt.Errorf("Field Md5sum no set on Md5sum %+v", self)

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		return *self.TargetFolderRelativeToTask, nil
		return nil, fmt.Errorf("Field TargetFolderRelativeToTask no set on TargetFolderRelativeToTask %+v", self)

	}
}

func (self *ExternalArtifact) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExternalArtifact", name)

	case "name", "Name":
		self.Name = nil

	case "url", "Url":
		self.Url = nil

	case "filesize", "Filesize":
		self.Filesize = nil

	case "isArtifactList", "IsArtifactList":
		self.IsArtifactList = nil

	case "filename", "Filename":
		self.Filename = nil

	case "md5sum", "Md5sum":
		self.Md5sum = nil

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		self.TargetFolderRelativeToTask = nil

	}

	return nil
}

func (self *ExternalArtifact) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ExternalArtifactList []*ExternalArtifact

func (self *ExternalArtifactList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ExternalArtifactList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ExternalArtifactList cannot copy the values from %#v", other)
}

func (list *ExternalArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
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
	return swaggering.FormatJSON(list)
}
