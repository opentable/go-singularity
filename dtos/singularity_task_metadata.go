package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskMetadataMetadataLevel string

const (
	SingularityTaskMetadataMetadataLevelINFO  SingularityTaskMetadataMetadataLevel = "INFO"
	SingularityTaskMetadataMetadataLevelWARN  SingularityTaskMetadataMetadataLevel = "WARN"
	SingularityTaskMetadataMetadataLevelERROR SingularityTaskMetadataMetadataLevel = "ERROR"
)

type SingularityTaskMetadata struct {
	Message   *string                               `json:"message,omitempty"`
	User      *string                               `json:"user,omitempty"`
	TaskId    *SingularityTaskId                    `json:"taskId,omitempty"`
	Timestamp *int64                                `json:"timestamp,omitempty"`
	Type      *string                               `json:"type,omitempty"`
	Title     *string                               `json:"title,omitempty"`
	Level     *SingularityTaskMetadataMetadataLevel `json:"level,omitempty"`
}

func (self *SingularityTaskMetadata) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskMetadata) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskMetadata); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskMetadata cannot copy the values from %#v", other)
}

func (self *SingularityTaskMetadata) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskMetadata) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskMetadata) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskMetadata", name)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			return nil
		}
		return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

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

	case "level", "Level":
		v, ok := value.(SingularityTaskMetadataMetadataLevel)
		if ok {
			self.Level = &v
			return nil
		}
		return fmt.Errorf("Field level/Level: value %v(%T) couldn't be cast to type SingularityTaskMetadataMetadataLevel", value, value)

	}
}

func (self *SingularityTaskMetadata) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskMetadata", name)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "taskId", "TaskId":
		return self.TaskId, nil
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "type", "Type":
		return *self.Type, nil
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "title", "Title":
		return *self.Title, nil
		return nil, fmt.Errorf("Field Title no set on Title %+v", self)

	case "level", "Level":
		return *self.Level, nil
		return nil, fmt.Errorf("Field Level no set on Level %+v", self)

	}
}

func (self *SingularityTaskMetadata) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskMetadata", name)

	case "message", "Message":
		self.Message = nil

	case "user", "User":
		self.User = nil

	case "taskId", "TaskId":
		self.TaskId = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "type", "Type":
		self.Type = nil

	case "title", "Title":
		self.Title = nil

	case "level", "Level":
		self.Level = nil

	}

	return nil
}

func (self *SingularityTaskMetadata) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskMetadataList []*SingularityTaskMetadata

func (self *SingularityTaskMetadataList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskMetadataList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskMetadataList cannot copy the values from %#v", other)
}

func (list *SingularityTaskMetadataList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskMetadataList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskMetadataList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
