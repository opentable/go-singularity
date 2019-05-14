package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityExpiringSkipHealthchecks struct {
	User *string `json:"user,omitempty"`
	// Invalid field: ExpiringAPIRequestObject *notfound.T `json:"expiringAPIRequestObject,omitempty"`
	RevertToSkipHealthchecks *bool   `json:"revertToSkipHealthchecks,omitempty"`
	RequestId                *string `json:"requestId,omitempty"`
	StartMillis              *int64  `json:"startMillis,omitempty"`
	ActionId                 *string `json:"actionId,omitempty"`
}

func (self *SingularityExpiringSkipHealthchecks) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityExpiringSkipHealthchecks) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExpiringSkipHealthchecks); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExpiringSkipHealthchecks cannot copy the values from %#v", other)
}

func (self *SingularityExpiringSkipHealthchecks) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityExpiringSkipHealthchecks) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityExpiringSkipHealthchecks) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringSkipHealthchecks", name)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "revertToSkipHealthchecks", "RevertToSkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.RevertToSkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field revertToSkipHealthchecks/RevertToSkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = &v
			return nil
		}
		return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)

	case "startMillis", "StartMillis":
		v, ok := value.(int64)
		if ok {
			self.StartMillis = &v
			return nil
		}
		return fmt.Errorf("Field startMillis/StartMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityExpiringSkipHealthchecks) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityExpiringSkipHealthchecks", name)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "revertToSkipHealthchecks", "RevertToSkipHealthchecks":
		return *self.RevertToSkipHealthchecks, nil
		return nil, fmt.Errorf("Field RevertToSkipHealthchecks no set on RevertToSkipHealthchecks %+v", self)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "startMillis", "StartMillis":
		return *self.StartMillis, nil
		return nil, fmt.Errorf("Field StartMillis no set on StartMillis %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	}
}

func (self *SingularityExpiringSkipHealthchecks) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringSkipHealthchecks", name)

	case "user", "User":
		self.User = nil

	case "revertToSkipHealthchecks", "RevertToSkipHealthchecks":
		self.RevertToSkipHealthchecks = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	case "startMillis", "StartMillis":
		self.StartMillis = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	}

	return nil
}

func (self *SingularityExpiringSkipHealthchecks) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityExpiringSkipHealthchecksList []*SingularityExpiringSkipHealthchecks

func (self *SingularityExpiringSkipHealthchecksList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExpiringSkipHealthchecksList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExpiringSkipHealthchecksList cannot copy the values from %#v", other)
}

func (list *SingularityExpiringSkipHealthchecksList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringSkipHealthchecksList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringSkipHealthchecksList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
