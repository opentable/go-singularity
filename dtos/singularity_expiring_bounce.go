package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityExpiringBounce struct {
	// Invalid field: ExpiringAPIRequestObject *notfound.T `json:"expiringAPIRequestObject,omitempty"`
	DeployId    *string `json:"deployId,omitempty"`
	RequestId   *string `json:"requestId,omitempty"`
	StartMillis *int64  `json:"startMillis,omitempty"`
	ActionId    *string `json:"actionId,omitempty"`
	User        *string `json:"user,omitempty"`
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

func (self *SingularityExpiringBounce) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityExpiringBounce) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityExpiringBounce) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringBounce", name)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = &v
			return nil
		}
		return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)

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

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityExpiringBounce) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityExpiringBounce", name)

	case "deployId", "DeployId":
		return *self.DeployId, nil
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "startMillis", "StartMillis":
		return *self.StartMillis, nil
		return nil, fmt.Errorf("Field StartMillis no set on StartMillis %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityExpiringBounce) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringBounce", name)

	case "deployId", "DeployId":
		self.DeployId = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	case "startMillis", "StartMillis":
		self.StartMillis = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "user", "User":
		self.User = nil

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
