package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type EmbeddedArtifact struct {
	present map[string]bool

	Content swaggering.StringList `json:"content"`

	Filename string `json:"filename,omitempty"`

	Md5sum string `json:"md5sum,omitempty"`

	Name string `json:"name,omitempty"`

	TargetFolderRelativeToTask string `json:"targetFolderRelativeToTask,omitempty"`
}

func (self *EmbeddedArtifact) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *EmbeddedArtifact) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*EmbeddedArtifact); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A EmbeddedArtifact cannot copy the values from %#v", other)
}

func (self *EmbeddedArtifact) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *EmbeddedArtifact) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *EmbeddedArtifact) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *EmbeddedArtifact) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *EmbeddedArtifact) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on EmbeddedArtifact", name)

	case "content", "Content":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Content = v
			self.present["content"] = true
			return nil
		} else {
			return fmt.Errorf("Field content/Content: value %v(%T) couldn't be cast to type StringList", value, value)
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

	}
}

func (self *EmbeddedArtifact) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on EmbeddedArtifact", name)

	case "content", "Content":
		if self.present != nil {
			if _, ok := self.present["content"]; ok {
				return self.Content, nil
			}
		}
		return nil, fmt.Errorf("Field Content no set on Content %+v", self)

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

	}
}

func (self *EmbeddedArtifact) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on EmbeddedArtifact", name)

	case "content", "Content":
		self.present["content"] = false

	case "filename", "Filename":
		self.present["filename"] = false

	case "md5sum", "Md5sum":
		self.present["md5sum"] = false

	case "name", "Name":
		self.present["name"] = false

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		self.present["targetFolderRelativeToTask"] = false

	}

	return nil
}

func (self *EmbeddedArtifact) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type EmbeddedArtifactList []*EmbeddedArtifact

func (self *EmbeddedArtifactList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*EmbeddedArtifactList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A EmbeddedArtifactList cannot copy the values from %#v", other)
}

func (list *EmbeddedArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *EmbeddedArtifactList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *EmbeddedArtifactList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
