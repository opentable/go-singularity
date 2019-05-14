package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityExpiringScale struct {
	Bounce      *bool   `json:"bounce,omitempty"`
	RequestId   *string `json:"requestId,omitempty"`
	StartMillis *int64  `json:"startMillis,omitempty"`
	ActionId    *string `json:"actionId,omitempty"`
	User        *string `json:"user,omitempty"`
	// Invalid field: ExpiringAPIRequestObject *notfound.T `json:"expiringAPIRequestObject,omitempty"`
	RevertToInstances *int32 `json:"revertToInstances,omitempty"`
}

func (self *SingularityExpiringScale) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityExpiringScale) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExpiringScale); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExpiringScale cannot copy the values from %#v", other)
}

func (self *SingularityExpiringScale) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityExpiringScale) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityExpiringScale) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringScale", name)

	case "bounce", "Bounce":
		v, ok := value.(bool)
		if ok {
			self.Bounce = &v
			return nil
		}
		return fmt.Errorf("Field bounce/Bounce: value %v(%T) couldn't be cast to type bool", value, value)

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

	case "revertToInstances", "RevertToInstances":
		v, ok := value.(int32)
		if ok {
			self.RevertToInstances = &v
			return nil
		}
		return fmt.Errorf("Field revertToInstances/RevertToInstances: value %v(%T) couldn't be cast to type int32", value, value)

	}
}

func (self *SingularityExpiringScale) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityExpiringScale", name)

	case "bounce", "Bounce":
		return *self.Bounce, nil
		return nil, fmt.Errorf("Field Bounce no set on Bounce %+v", self)

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

	case "revertToInstances", "RevertToInstances":
		return *self.RevertToInstances, nil
		return nil, fmt.Errorf("Field RevertToInstances no set on RevertToInstances %+v", self)

	}
}

func (self *SingularityExpiringScale) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringScale", name)

	case "bounce", "Bounce":
		self.Bounce = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	case "startMillis", "StartMillis":
		self.StartMillis = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "user", "User":
		self.User = nil

	case "revertToInstances", "RevertToInstances":
		self.RevertToInstances = nil

	}

	return nil
}

func (self *SingularityExpiringScale) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityExpiringScaleList []*SingularityExpiringScale

func (self *SingularityExpiringScaleList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExpiringScaleList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExpiringScaleList cannot copy the values from %#v", other)
}

func (list *SingularityExpiringScaleList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringScaleList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringScaleList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
