package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskMetadataRequestMetadataLevel string

const (
	SingularityTaskMetadataRequestMetadataLevelINFO  SingularityTaskMetadataRequestMetadataLevel = "INFO"
	SingularityTaskMetadataRequestMetadataLevelWARN  SingularityTaskMetadataRequestMetadataLevel = "WARN"
	SingularityTaskMetadataRequestMetadataLevelERROR SingularityTaskMetadataRequestMetadataLevel = "ERROR"
)

type SingularityTaskMetadataRequest struct {
	Type    *string                                      `json:"type,omitempty"`
	Title   *string                                      `json:"title,omitempty"`
	Message *string                                      `json:"message,omitempty"`
	Level   *SingularityTaskMetadataRequestMetadataLevel `json:"level,omitempty"`
}

func (self *SingularityTaskMetadataRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskMetadataRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskMetadataRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskMetadataRequest cannot copy the values from %#v", other)
}

func (self *SingularityTaskMetadataRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskMetadataRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskMetadataRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskMetadataRequest", name)

	case "type", "Type":
		v, ok := value.(string)
		if ok {
			self.Type = &v
			return nil
		}
		return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type string", value, value)

	case "title", "Title":
		v, ok := value.(string)
		if ok {
			self.Title = &v
			return nil
		}
		return fmt.Errorf("Field title/Title: value %v(%T) couldn't be cast to type string", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "level", "Level":
		v, ok := value.(SingularityTaskMetadataRequestMetadataLevel)
		if ok {
			self.Level = &v
			return nil
		}
		return fmt.Errorf("Field level/Level: value %v(%T) couldn't be cast to type SingularityTaskMetadataRequestMetadataLevel", value, value)

	}
}

func (self *SingularityTaskMetadataRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskMetadataRequest", name)

	case "type", "Type":
		return *self.Type, nil
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "title", "Title":
		return *self.Title, nil
		return nil, fmt.Errorf("Field Title no set on Title %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "level", "Level":
		return *self.Level, nil
		return nil, fmt.Errorf("Field Level no set on Level %+v", self)

	}
}

func (self *SingularityTaskMetadataRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskMetadataRequest", name)

	case "type", "Type":
		self.Type = nil

	case "title", "Title":
		self.Title = nil

	case "message", "Message":
		self.Message = nil

	case "level", "Level":
		self.Level = nil

	}

	return nil
}

func (self *SingularityTaskMetadataRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskMetadataRequestList []*SingularityTaskMetadataRequest

func (self *SingularityTaskMetadataRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskMetadataRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskMetadataRequestList cannot copy the values from %#v", other)
}

func (list *SingularityTaskMetadataRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskMetadataRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskMetadataRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
