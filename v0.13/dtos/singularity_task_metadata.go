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
	present map[string]bool

	Level SingularityTaskMetadataMetadataLevel `json:"level"`

	Message string `json:"message,omitempty"`

	TaskId *SingularityTaskId `json:"taskId"`

	Timestamp int64 `json:"timestamp"`

	Title string `json:"title,omitempty"`

	Type string `json:"type,omitempty"`

	User string `json:"user,omitempty"`
}

func (self *SingularityTaskMetadata) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskMetadata) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskMetadata); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskMetadata cannot absorb the values from %v", other)
}

func (self *SingularityTaskMetadata) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskMetadata) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskMetadata) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskMetadata) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskMetadata) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskMetadata", name)

	case "level", "Level":
		v, ok := value.(SingularityTaskMetadataMetadataLevel)
		if ok {
			self.Level = v
			self.present["level"] = true
			return nil
		} else {
			return fmt.Errorf("Field level/Level: value %v(%T) couldn't be cast to type SingularityTaskMetadataMetadataLevel", value, value)
		}

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			self.present["taskId"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)
		}

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "title", "Title":
		v, ok := value.(string)
		if ok {
			self.Title = v
			self.present["title"] = true
			return nil
		} else {
			return fmt.Errorf("Field title/Title: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "type", "Type":
		v, ok := value.(string)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityTaskMetadata) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskMetadata", name)

	case "level", "Level":
		if self.present != nil {
			if _, ok := self.present["level"]; ok {
				return self.Level, nil
			}
		}
		return nil, fmt.Errorf("Field Level no set on Level %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "taskId", "TaskId":
		if self.present != nil {
			if _, ok := self.present["taskId"]; ok {
				return self.TaskId, nil
			}
		}
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "title", "Title":
		if self.present != nil {
			if _, ok := self.present["title"]; ok {
				return self.Title, nil
			}
		}
		return nil, fmt.Errorf("Field Title no set on Title %+v", self)

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityTaskMetadata) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskMetadata", name)

	case "level", "Level":
		self.present["level"] = false

	case "message", "Message":
		self.present["message"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	case "title", "Title":
		self.present["title"] = false

	case "type", "Type":
		self.present["type"] = false

	case "user", "User":
		self.present["user"] = false

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
	return fmt.Errorf("A SingularityTaskMetadata cannot absorb the values from %v", other)
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
