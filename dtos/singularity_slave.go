package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularitySlave struct {
	Host         *string                               `json:"host,omitempty"`
	RackId       *string                               `json:"rackId,omitempty"`
	Attributes   *map[string]string                    `json:"attributes,omitempty"`
	Resources    *MesosResourcesObject                 `json:"resources,omitempty"`
	CurrentState *SingularityMachineStateHistoryUpdate `json:"currentState,omitempty"`
	FirstSeenAt  *int64                                `json:"firstSeenAt,omitempty"`
	Id           *string                               `json:"id,omitempty"`
}

func (self *SingularitySlave) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularitySlave) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularitySlave); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularitySlave cannot copy the values from %#v", other)
}

func (self *SingularitySlave) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularitySlave) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularitySlave) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySlave", name)

	case "host", "Host":
		v, ok := value.(string)
		if ok {
			self.Host = &v
			return nil
		}
		return fmt.Errorf("Field host/Host: value %v(%T) couldn't be cast to type string", value, value)

	case "rackId", "RackId":
		v, ok := value.(string)
		if ok {
			self.RackId = &v
			return nil
		}
		return fmt.Errorf("Field rackId/RackId: value %v(%T) couldn't be cast to type string", value, value)

	case "attributes", "Attributes":
		v, ok := value.(map[string]string)
		if ok {
			self.Attributes = &v
			return nil
		}
		return fmt.Errorf("Field attributes/Attributes: value %v(%T) couldn't be cast to type map[string]string", value, value)

	case "resources", "Resources":
		v, ok := value.(*MesosResourcesObject)
		if ok {
			self.Resources = v
			return nil
		}
		return fmt.Errorf("Field resources/Resources: value %v(%T) couldn't be cast to type *MesosResourcesObject", value, value)

	case "currentState", "CurrentState":
		v, ok := value.(*SingularityMachineStateHistoryUpdate)
		if ok {
			self.CurrentState = v
			return nil
		}
		return fmt.Errorf("Field currentState/CurrentState: value %v(%T) couldn't be cast to type *SingularityMachineStateHistoryUpdate", value, value)

	case "firstSeenAt", "FirstSeenAt":
		v, ok := value.(int64)
		if ok {
			self.FirstSeenAt = &v
			return nil
		}
		return fmt.Errorf("Field firstSeenAt/FirstSeenAt: value %v(%T) couldn't be cast to type int64", value, value)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = &v
			return nil
		}
		return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularitySlave) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularitySlave", name)

	case "host", "Host":
		return *self.Host, nil
		return nil, fmt.Errorf("Field Host no set on Host %+v", self)

	case "rackId", "RackId":
		return *self.RackId, nil
		return nil, fmt.Errorf("Field RackId no set on RackId %+v", self)

	case "attributes", "Attributes":
		return *self.Attributes, nil
		return nil, fmt.Errorf("Field Attributes no set on Attributes %+v", self)

	case "resources", "Resources":
		return self.Resources, nil
		return nil, fmt.Errorf("Field Resources no set on Resources %+v", self)

	case "currentState", "CurrentState":
		return self.CurrentState, nil
		return nil, fmt.Errorf("Field CurrentState no set on CurrentState %+v", self)

	case "firstSeenAt", "FirstSeenAt":
		return *self.FirstSeenAt, nil
		return nil, fmt.Errorf("Field FirstSeenAt no set on FirstSeenAt %+v", self)

	case "id", "Id":
		return *self.Id, nil
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	}
}

func (self *SingularitySlave) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySlave", name)

	case "host", "Host":
		self.Host = nil

	case "rackId", "RackId":
		self.RackId = nil

	case "attributes", "Attributes":
		self.Attributes = nil

	case "resources", "Resources":
		self.Resources = nil

	case "currentState", "CurrentState":
		self.CurrentState = nil

	case "firstSeenAt", "FirstSeenAt":
		self.FirstSeenAt = nil

	case "id", "Id":
		self.Id = nil

	}

	return nil
}

func (self *SingularitySlave) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularitySlaveList []*SingularitySlave

func (self *SingularitySlaveList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularitySlaveList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularitySlaveList cannot copy the values from %#v", other)
}

func (list *SingularitySlaveList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularitySlaveList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularitySlaveList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
