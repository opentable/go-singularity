package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeployMarker struct {
	DeployId  *string `json:"deployId,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty"`
	User      *string `json:"user,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
}

func (self *SingularityDeployMarker) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployMarker) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployMarker); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployMarker cannot copy the values from %#v", other)
}

func (self *SingularityDeployMarker) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployMarker) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployMarker) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployMarker", name)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = &v
			return nil
		}
		return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = &v
			return nil
		}
		return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityDeployMarker) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployMarker", name)

	case "deployId", "DeployId":
		return *self.DeployId, nil
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	}
}

func (self *SingularityDeployMarker) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployMarker", name)

	case "deployId", "DeployId":
		self.DeployId = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "user", "User":
		self.User = nil

	case "message", "Message":
		self.Message = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	}

	return nil
}

func (self *SingularityDeployMarker) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployMarkerList []*SingularityDeployMarker

func (self *SingularityDeployMarkerList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployMarkerList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployMarkerList cannot copy the values from %#v", other)
}

func (list *SingularityDeployMarkerList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployMarkerList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployMarkerList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
