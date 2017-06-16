package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ExternalArtifact struct {
	present map[string]bool

	Filename string `json:"filename,omitempty"`

	Filesize int64 `json:"filesize"`

	Md5sum string `json:"md5sum,omitempty"`

	Name string `json:"name,omitempty"`

	TargetFolderRelativeToTask string `json:"targetFolderRelativeToTask,omitempty"`

	Url string `json:"url,omitempty"`
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

func (self *ExternalArtifact) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *ExternalArtifact) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ExternalArtifact) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ExternalArtifact) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *ExternalArtifact) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExternalArtifact", name)

	case "filename", "Filename":
		v, ok := value.(string)
		if ok {
			self.Filename = v
			self.present["filename"] = true
			return nil
		} else {
			return fmt.Errorf("Field filename/Filename: value %v(%T) couldn't be cast to type string", value, value)
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

	case "md5sum", "Md5sum":
		v, ok := value.(string)
		if ok {
			self.Md5sum = v
			self.present["md5sum"] = true
			return nil
		} else {
			return fmt.Errorf("Field md5sum/Md5sum: value %v(%T) couldn't be cast to type string", value, value)
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

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		v, ok := value.(string)
		if ok {
			self.TargetFolderRelativeToTask = v
			self.present["targetFolderRelativeToTask"] = true
			return nil
		} else {
			return fmt.Errorf("Field targetFolderRelativeToTask/TargetFolderRelativeToTask: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "url", "Url":
		v, ok := value.(string)
		if ok {
			self.Url = v
			self.present["url"] = true
			return nil
		} else {
			return fmt.Errorf("Field url/Url: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *ExternalArtifact) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ExternalArtifact", name)

	case "filename", "Filename":
		if self.present != nil {
			if _, ok := self.present["filename"]; ok {
				return self.Filename, nil
			}
		}
		return nil, fmt.Errorf("Field Filename no set on Filename %+v", self)

	case "filesize", "Filesize":
		if self.present != nil {
			if _, ok := self.present["filesize"]; ok {
				return self.Filesize, nil
			}
		}
		return nil, fmt.Errorf("Field Filesize no set on Filesize %+v", self)

	case "md5sum", "Md5sum":
		if self.present != nil {
			if _, ok := self.present["md5sum"]; ok {
				return self.Md5sum, nil
			}
		}
		return nil, fmt.Errorf("Field Md5sum no set on Md5sum %+v", self)

	case "name", "Name":
		if self.present != nil {
			if _, ok := self.present["name"]; ok {
				return self.Name, nil
			}
		}
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		if self.present != nil {
			if _, ok := self.present["targetFolderRelativeToTask"]; ok {
				return self.TargetFolderRelativeToTask, nil
			}
		}
		return nil, fmt.Errorf("Field TargetFolderRelativeToTask no set on TargetFolderRelativeToTask %+v", self)

	case "url", "Url":
		if self.present != nil {
			if _, ok := self.present["url"]; ok {
				return self.Url, nil
			}
		}
		return nil, fmt.Errorf("Field Url no set on Url %+v", self)

	}
}

func (self *ExternalArtifact) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExternalArtifact", name)

	case "filename", "Filename":
		self.present["filename"] = false

	case "filesize", "Filesize":
		self.present["filesize"] = false

	case "md5sum", "Md5sum":
		self.present["md5sum"] = false

	case "name", "Name":
		self.present["name"] = false

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		self.present["targetFolderRelativeToTask"] = false

	case "url", "Url":
		self.present["url"] = false

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
