package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityWebhookWebhookType string

const (
	SingularityWebhookWebhookTypeTASK    SingularityWebhookWebhookType = "TASK"
	SingularityWebhookWebhookTypeREQUEST SingularityWebhookWebhookType = "REQUEST"
	SingularityWebhookWebhookTypeDEPLOY  SingularityWebhookWebhookType = "DEPLOY"
)

type SingularityWebhook struct {
	present map[string]bool

	Id string `json:"id,omitempty"`

	Timestamp int64 `json:"timestamp"`

	Type SingularityWebhookWebhookType `json:"type"`

	Uri string `json:"uri,omitempty"`

	User string `json:"user,omitempty"`
}

func (self *SingularityWebhook) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityWebhook) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityWebhook); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityWebhook cannot copy the values from %#v", other)
}

func (self *SingularityWebhook) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityWebhook) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityWebhook) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityWebhook) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityWebhook) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityWebhook", name)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
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

	case "type", "Type":
		v, ok := value.(SingularityWebhookWebhookType)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type SingularityWebhookWebhookType", value, value)
		}

	case "uri", "Uri":
		v, ok := value.(string)
		if ok {
			self.Uri = v
			self.present["uri"] = true
			return nil
		} else {
			return fmt.Errorf("Field uri/Uri: value %v(%T) couldn't be cast to type string", value, value)
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

func (self *SingularityWebhook) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityWebhook", name)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "uri", "Uri":
		if self.present != nil {
			if _, ok := self.present["uri"]; ok {
				return self.Uri, nil
			}
		}
		return nil, fmt.Errorf("Field Uri no set on Uri %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityWebhook) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityWebhook", name)

	case "id", "Id":
		self.present["id"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	case "type", "Type":
		self.present["type"] = false

	case "uri", "Uri":
		self.present["uri"] = false

	case "user", "User":
		self.present["user"] = false

	}

	return nil
}

func (self *SingularityWebhook) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityWebhookList []*SingularityWebhook

func (self *SingularityWebhookList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityWebhookList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityWebhookList cannot copy the values from %#v", other)
}

func (list *SingularityWebhookList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityWebhookList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityWebhookList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
