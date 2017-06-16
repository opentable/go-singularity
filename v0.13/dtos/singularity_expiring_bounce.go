package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityExpiringBounce struct {
	present map[string]bool

	ActionId string `json:"actionId,omitempty"`

	DeployId string `json:"deployId,omitempty"`

	// ExpiringAPIRequestObject *T `json:"expiringAPIRequestObject"`

	RequestId string `json:"requestId,omitempty"`

	StartMillis int64 `json:"startMillis"`

	User string `json:"user,omitempty"`
}

func (self *SingularityExpiringBounce) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityExpiringBounce) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExpiringBounce); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExpiringBounce cannot copy the values from %#v", other)
}

func (self *SingularityExpiringBounce) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityExpiringBounce) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityExpiringBounce) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityExpiringBounce) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityExpiringBounce) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringBounce", name)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = v
			self.present["deployId"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = v
			self.present["requestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "startMillis", "StartMillis":
		v, ok := value.(int64)
		if ok {
			self.StartMillis = v
			self.present["startMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field startMillis/StartMillis: value %v(%T) couldn't be cast to type int64", value, value)
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

func (self *SingularityExpiringBounce) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityExpiringBounce", name)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "deployId", "DeployId":
		if self.present != nil {
			if _, ok := self.present["deployId"]; ok {
				return self.DeployId, nil
			}
		}
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "startMillis", "StartMillis":
		if self.present != nil {
			if _, ok := self.present["startMillis"]; ok {
				return self.StartMillis, nil
			}
		}
		return nil, fmt.Errorf("Field StartMillis no set on StartMillis %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityExpiringBounce) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringBounce", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "deployId", "DeployId":
		self.present["deployId"] = false

	case "requestId", "RequestId":
		self.present["requestId"] = false

	case "startMillis", "StartMillis":
		self.present["startMillis"] = false

	case "user", "User":
		self.present["user"] = false

	}

	return nil
}

func (self *SingularityExpiringBounce) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityExpiringBounceList []*SingularityExpiringBounce

func (self *SingularityExpiringBounceList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExpiringBounceList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExpiringBounceList cannot copy the values from %#v", other)
}

func (list *SingularityExpiringBounceList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringBounceList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringBounceList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
