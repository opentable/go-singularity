package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskId struct {
	present map[string]bool

	DeployId string `json:"deployId,omitempty"`

	Host string `json:"host,omitempty"`

	Id string `json:"id,omitempty"`

	InstanceNo int32 `json:"instanceNo"`

	RackId string `json:"rackId,omitempty"`

	RequestId string `json:"requestId,omitempty"`

	SanitizedHost string `json:"sanitizedHost,omitempty"`

	SanitizedRackId string `json:"sanitizedRackId,omitempty"`

	StartedAt int64 `json:"startedAt"`
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

func (self *SingularityTaskId) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskId) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskId) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskId) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskId) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskId", name)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = v
			self.present["deployId"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "host", "Host":
		v, ok := value.(string)
		if ok {
			self.Host = v
			self.present["host"] = true
			return nil
		} else {
			return fmt.Errorf("Field host/Host: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "instanceNo", "InstanceNo":
		v, ok := value.(int32)
		if ok {
			self.InstanceNo = v
			self.present["instanceNo"] = true
			return nil
		} else {
			return fmt.Errorf("Field instanceNo/InstanceNo: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "rackId", "RackId":
		v, ok := value.(string)
		if ok {
			self.RackId = v
			self.present["rackId"] = true
			return nil
		} else {
			return fmt.Errorf("Field rackId/RackId: value %v(%T) couldn't be cast to type string", value, value)
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

	case "sanitizedHost", "SanitizedHost":
		v, ok := value.(string)
		if ok {
			self.SanitizedHost = v
			self.present["sanitizedHost"] = true
			return nil
		} else {
			return fmt.Errorf("Field sanitizedHost/SanitizedHost: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "sanitizedRackId", "SanitizedRackId":
		v, ok := value.(string)
		if ok {
			self.SanitizedRackId = v
			self.present["sanitizedRackId"] = true
			return nil
		} else {
			return fmt.Errorf("Field sanitizedRackId/SanitizedRackId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "startedAt", "StartedAt":
		v, ok := value.(int64)
		if ok {
			self.StartedAt = v
			self.present["startedAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field startedAt/StartedAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityTaskId) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskId", name)

	case "deployId", "DeployId":
		if self.present != nil {
			if _, ok := self.present["deployId"]; ok {
				return self.DeployId, nil
			}
		}
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "host", "Host":
		if self.present != nil {
			if _, ok := self.present["host"]; ok {
				return self.Host, nil
			}
		}
		return nil, fmt.Errorf("Field Host no set on Host %+v", self)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "instanceNo", "InstanceNo":
		if self.present != nil {
			if _, ok := self.present["instanceNo"]; ok {
				return self.InstanceNo, nil
			}
		}
		return nil, fmt.Errorf("Field InstanceNo no set on InstanceNo %+v", self)

	case "rackId", "RackId":
		if self.present != nil {
			if _, ok := self.present["rackId"]; ok {
				return self.RackId, nil
			}
		}
		return nil, fmt.Errorf("Field RackId no set on RackId %+v", self)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "sanitizedHost", "SanitizedHost":
		if self.present != nil {
			if _, ok := self.present["sanitizedHost"]; ok {
				return self.SanitizedHost, nil
			}
		}
		return nil, fmt.Errorf("Field SanitizedHost no set on SanitizedHost %+v", self)

	case "sanitizedRackId", "SanitizedRackId":
		if self.present != nil {
			if _, ok := self.present["sanitizedRackId"]; ok {
				return self.SanitizedRackId, nil
			}
		}
		return nil, fmt.Errorf("Field SanitizedRackId no set on SanitizedRackId %+v", self)

	case "startedAt", "StartedAt":
		if self.present != nil {
			if _, ok := self.present["startedAt"]; ok {
				return self.StartedAt, nil
			}
		}
		return nil, fmt.Errorf("Field StartedAt no set on StartedAt %+v", self)

	}
}

func (self *SingularityTaskId) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskId", name)

	case "deployId", "DeployId":
		self.present["deployId"] = false

	case "host", "Host":
		self.present["host"] = false

	case "id", "Id":
		self.present["id"] = false

	case "instanceNo", "InstanceNo":
		self.present["instanceNo"] = false

	case "rackId", "RackId":
		self.present["rackId"] = false

	case "requestId", "RequestId":
		self.present["requestId"] = false

	case "sanitizedHost", "SanitizedHost":
		self.present["sanitizedHost"] = false

	case "sanitizedRackId", "SanitizedRackId":
		self.present["sanitizedRackId"] = false

	case "startedAt", "StartedAt":
		self.present["startedAt"] = false

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
