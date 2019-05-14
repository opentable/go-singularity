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
	Uri       *string                        `json:"uri,omitempty"`
	Type      *SingularityWebhookWebhookType `json:"type,omitempty"`
	User      *string                        `json:"user,omitempty"`
	Timestamp *int64                         `json:"timestamp,omitempty"`
	Id        *string                        `json:"id,omitempty"`
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

func (self *SingularityWebhook) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityWebhook) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityWebhook) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityWebhook", name)

	case "uri", "Uri":
		v, ok := value.(string)
		if ok {
			self.Uri = &v
			return nil
		}
		return fmt.Errorf("Field uri/Uri: value %v(%T) couldn't be cast to type string", value, value)

	case "type", "Type":
		v, ok := value.(SingularityWebhookWebhookType)
		if ok {
			self.Type = &v
			return nil
		}
		return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type SingularityWebhookWebhookType", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = &v
			return nil
		}
		return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityWebhook) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityWebhook", name)

	case "uri", "Uri":
		return *self.Uri, nil
		return nil, fmt.Errorf("Field Uri no set on Uri %+v", self)

	case "type", "Type":
		return *self.Type, nil
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "id", "Id":
		return *self.Id, nil
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	}
}

func (self *SingularityWebhook) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityWebhook", name)

	case "uri", "Uri":
		self.Uri = nil

	case "type", "Type":
		self.Type = nil

	case "user", "User":
		self.User = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "id", "Id":
		self.Id = nil

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
