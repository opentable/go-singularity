package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskId struct {
	RequestId       *string `json:"requestId,omitempty"`
	StartedAt       *int64  `json:"startedAt,omitempty"`
	InstanceNo      *int32  `json:"instanceNo,omitempty"`
	SanitizedHost   *string `json:"sanitizedHost,omitempty"`
	RackId          *string `json:"rackId,omitempty"`
	DeployId        *string `json:"deployId,omitempty"`
	SanitizedRackId *string `json:"sanitizedRackId,omitempty"`
	Host            *string `json:"host,omitempty"`
	Id              *string `json:"id,omitempty"`
}

func (self *SingularityTaskId) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskId) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskId); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskId cannot copy the values from %#v", other)
}

func (self *SingularityTaskId) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskId) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskId) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskId", name)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = &v
			return nil
		}
		return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)

	case "startedAt", "StartedAt":
		v, ok := value.(int64)
		if ok {
			self.StartedAt = &v
			return nil
		}
		return fmt.Errorf("Field startedAt/StartedAt: value %v(%T) couldn't be cast to type int64", value, value)

	case "instanceNo", "InstanceNo":
		v, ok := value.(int32)
		if ok {
			self.InstanceNo = &v
			return nil
		}
		return fmt.Errorf("Field instanceNo/InstanceNo: value %v(%T) couldn't be cast to type int32", value, value)

	case "sanitizedHost", "SanitizedHost":
		v, ok := value.(string)
		if ok {
			self.SanitizedHost = &v
			return nil
		}
		return fmt.Errorf("Field sanitizedHost/SanitizedHost: value %v(%T) couldn't be cast to type string", value, value)

	case "rackId", "RackId":
		v, ok := value.(string)
		if ok {
			self.RackId = &v
			return nil
		}
		return fmt.Errorf("Field rackId/RackId: value %v(%T) couldn't be cast to type string", value, value)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = &v
			return nil
		}
		return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)

	case "sanitizedRackId", "SanitizedRackId":
		v, ok := value.(string)
		if ok {
			self.SanitizedRackId = &v
			return nil
		}
		return fmt.Errorf("Field sanitizedRackId/SanitizedRackId: value %v(%T) couldn't be cast to type string", value, value)

	case "host", "Host":
		v, ok := value.(string)
		if ok {
			self.Host = &v
			return nil
		}
		return fmt.Errorf("Field host/Host: value %v(%T) couldn't be cast to type string", value, value)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = &v
			return nil
		}
		return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityTaskId) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskId", name)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "startedAt", "StartedAt":
		return *self.StartedAt, nil
		return nil, fmt.Errorf("Field StartedAt no set on StartedAt %+v", self)

	case "instanceNo", "InstanceNo":
		return *self.InstanceNo, nil
		return nil, fmt.Errorf("Field InstanceNo no set on InstanceNo %+v", self)

	case "sanitizedHost", "SanitizedHost":
		return *self.SanitizedHost, nil
		return nil, fmt.Errorf("Field SanitizedHost no set on SanitizedHost %+v", self)

	case "rackId", "RackId":
		return *self.RackId, nil
		return nil, fmt.Errorf("Field RackId no set on RackId %+v", self)

	case "deployId", "DeployId":
		return *self.DeployId, nil
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "sanitizedRackId", "SanitizedRackId":
		return *self.SanitizedRackId, nil
		return nil, fmt.Errorf("Field SanitizedRackId no set on SanitizedRackId %+v", self)

	case "host", "Host":
		return *self.Host, nil
		return nil, fmt.Errorf("Field Host no set on Host %+v", self)

	case "id", "Id":
		return *self.Id, nil
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	}
}

func (self *SingularityTaskId) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskId", name)

	case "requestId", "RequestId":
		self.RequestId = nil

	case "startedAt", "StartedAt":
		self.StartedAt = nil

	case "instanceNo", "InstanceNo":
		self.InstanceNo = nil

	case "sanitizedHost", "SanitizedHost":
		self.SanitizedHost = nil

	case "rackId", "RackId":
		self.RackId = nil

	case "deployId", "DeployId":
		self.DeployId = nil

	case "sanitizedRackId", "SanitizedRackId":
		self.SanitizedRackId = nil

	case "host", "Host":
		self.Host = nil

	case "id", "Id":
		self.Id = nil

	}

	return nil
}

func (self *SingularityTaskId) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskIdList []*SingularityTaskId

func (self *SingularityTaskIdList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskIdList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskIdList cannot copy the values from %#v", other)
}

func (list *SingularityTaskIdList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskIdList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskIdList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
