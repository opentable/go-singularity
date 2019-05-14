package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type EmbeddedArtifact struct {
	Filename                   *string                `json:"filename,omitempty"`
	Md5sum                     *string                `json:"md5sum,omitempty"`
	TargetFolderRelativeToTask *string                `json:"targetFolderRelativeToTask,omitempty"`
	Name                       *string                `json:"name,omitempty"`
	Content                    *swaggering.StringList `json:"content,omitempty"`
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

func (self *EmbeddedArtifact) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *EmbeddedArtifact) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *EmbeddedArtifact) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on EmbeddedArtifact", name)

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

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = &v
			return nil
		}
		return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)

	case "content", "Content":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Content = &v
			return nil
		}
		return fmt.Errorf("Field content/Content: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	}
}

func (self *EmbeddedArtifact) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on EmbeddedArtifact", name)

	case "filename", "Filename":
		return *self.Filename, nil
		return nil, fmt.Errorf("Field Filename no set on Filename %+v", self)

	case "md5sum", "Md5sum":
		return *self.Md5sum, nil
		return nil, fmt.Errorf("Field Md5sum no set on Md5sum %+v", self)

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		return *self.TargetFolderRelativeToTask, nil
		return nil, fmt.Errorf("Field TargetFolderRelativeToTask no set on TargetFolderRelativeToTask %+v", self)

	case "name", "Name":
		return *self.Name, nil
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "content", "Content":
		return *self.Content, nil
		return nil, fmt.Errorf("Field Content no set on Content %+v", self)

	}
}

func (self *EmbeddedArtifact) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on EmbeddedArtifact", name)

	case "filename", "Filename":
		self.Filename = nil

	case "md5sum", "Md5sum":
		self.Md5sum = nil

	case "targetFolderRelativeToTask", "TargetFolderRelativeToTask":
		self.TargetFolderRelativeToTask = nil

	case "name", "Name":
		self.Name = nil

	case "content", "Content":
		self.Content = nil

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
