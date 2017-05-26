package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Offer struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	AttributesCount int32 `json:"attributesCount"`

	// AttributesList *List[Attribute] `json:"attributesList"`

	// AttributesOrBuilderList *List[? extends org.apache.mesos.Protos$AttributeOrBuilder] `json:"attributesOrBuilderList"`

	DefaultInstanceForType *Offer `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	ExecutorIdsCount int32 `json:"executorIdsCount"`

	ExecutorIdsList ExecutorIDList `json:"executorIdsList"`

	// ExecutorIdsOrBuilderList *List[? extends org.apache.mesos.Protos$ExecutorIDOrBuilder] `json:"executorIdsOrBuilderList"`

	FrameworkId *FrameworkID `json:"frameworkId"`

	FrameworkIdOrBuilder *FrameworkIDOrBuilder `json:"frameworkIdOrBuilder"`

	Hostname string `json:"hostname,omitempty"`

	HostnameBytes *ByteString `json:"hostnameBytes"`

	Id *OfferID `json:"id"`

	IdOrBuilder *OfferIDOrBuilder `json:"idOrBuilder"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$Offer> `json:"parserForType"`

	ResourcesCount int32 `json:"resourcesCount"`

	// ResourcesList *List[Resource] `json:"resourcesList"`

	// ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder] `json:"resourcesOrBuilderList"`

	SerializedSize int32 `json:"serializedSize"`

	SlaveId *SlaveID `json:"slaveId"`

	SlaveIdOrBuilder *SlaveIDOrBuilder `json:"slaveIdOrBuilder"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *Offer) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Offer) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Offer); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Offer cannot copy the values from %#v", other)
}

func (self *Offer) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Offer) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Offer) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Offer) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Offer) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Offer", name)

	case "attributesCount", "AttributesCount":
		v, ok := value.(int32)
		if ok {
			self.AttributesCount = v
			self.present["attributesCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field attributesCount/AttributesCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Offer)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *Offer", value, value)
		}

	case "descriptorForType", "DescriptorForType":
		v, ok := value.(*Descriptor)
		if ok {
			self.DescriptorForType = v
			self.present["descriptorForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field descriptorForType/DescriptorForType: value %v(%T) couldn't be cast to type *Descriptor", value, value)
		}

	case "executorIdsCount", "ExecutorIdsCount":
		v, ok := value.(int32)
		if ok {
			self.ExecutorIdsCount = v
			self.present["executorIdsCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field executorIdsCount/ExecutorIdsCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "executorIdsList", "ExecutorIdsList":
		v, ok := value.(ExecutorIDList)
		if ok {
			self.ExecutorIdsList = v
			self.present["executorIdsList"] = true
			return nil
		} else {
			return fmt.Errorf("Field executorIdsList/ExecutorIdsList: value %v(%T) couldn't be cast to type ExecutorIDList", value, value)
		}

	case "frameworkId", "FrameworkId":
		v, ok := value.(*FrameworkID)
		if ok {
			self.FrameworkId = v
			self.present["frameworkId"] = true
			return nil
		} else {
			return fmt.Errorf("Field frameworkId/FrameworkId: value %v(%T) couldn't be cast to type *FrameworkID", value, value)
		}

	case "frameworkIdOrBuilder", "FrameworkIdOrBuilder":
		v, ok := value.(*FrameworkIDOrBuilder)
		if ok {
			self.FrameworkIdOrBuilder = v
			self.present["frameworkIdOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field frameworkIdOrBuilder/FrameworkIdOrBuilder: value %v(%T) couldn't be cast to type *FrameworkIDOrBuilder", value, value)
		}

	case "hostname", "Hostname":
		v, ok := value.(string)
		if ok {
			self.Hostname = v
			self.present["hostname"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostname/Hostname: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "hostnameBytes", "HostnameBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.HostnameBytes = v
			self.present["hostnameBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostnameBytes/HostnameBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "id", "Id":
		v, ok := value.(*OfferID)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type *OfferID", value, value)
		}

	case "idOrBuilder", "IdOrBuilder":
		v, ok := value.(*OfferIDOrBuilder)
		if ok {
			self.IdOrBuilder = v
			self.present["idOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field idOrBuilder/IdOrBuilder: value %v(%T) couldn't be cast to type *OfferIDOrBuilder", value, value)
		}

	case "initializationErrorString", "InitializationErrorString":
		v, ok := value.(string)
		if ok {
			self.InitializationErrorString = v
			self.present["initializationErrorString"] = true
			return nil
		} else {
			return fmt.Errorf("Field initializationErrorString/InitializationErrorString: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "initialized", "Initialized":
		v, ok := value.(bool)
		if ok {
			self.Initialized = v
			self.present["initialized"] = true
			return nil
		} else {
			return fmt.Errorf("Field initialized/Initialized: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "resourcesCount", "ResourcesCount":
		v, ok := value.(int32)
		if ok {
			self.ResourcesCount = v
			self.present["resourcesCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field resourcesCount/ResourcesCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "slaveId", "SlaveId":
		v, ok := value.(*SlaveID)
		if ok {
			self.SlaveId = v
			self.present["slaveId"] = true
			return nil
		} else {
			return fmt.Errorf("Field slaveId/SlaveId: value %v(%T) couldn't be cast to type *SlaveID", value, value)
		}

	case "slaveIdOrBuilder", "SlaveIdOrBuilder":
		v, ok := value.(*SlaveIDOrBuilder)
		if ok {
			self.SlaveIdOrBuilder = v
			self.present["slaveIdOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field slaveIdOrBuilder/SlaveIdOrBuilder: value %v(%T) couldn't be cast to type *SlaveIDOrBuilder", value, value)
		}

	case "unknownFields", "UnknownFields":
		v, ok := value.(*UnknownFieldSet)
		if ok {
			self.UnknownFields = v
			self.present["unknownFields"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownFields/UnknownFields: value %v(%T) couldn't be cast to type *UnknownFieldSet", value, value)
		}

	}
}

func (self *Offer) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Offer", name)

	case "attributesCount", "AttributesCount":
		if self.present != nil {
			if _, ok := self.present["attributesCount"]; ok {
				return self.AttributesCount, nil
			}
		}
		return nil, fmt.Errorf("Field AttributesCount no set on AttributesCount %+v", self)

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

	case "descriptorForType", "DescriptorForType":
		if self.present != nil {
			if _, ok := self.present["descriptorForType"]; ok {
				return self.DescriptorForType, nil
			}
		}
		return nil, fmt.Errorf("Field DescriptorForType no set on DescriptorForType %+v", self)

	case "executorIdsCount", "ExecutorIdsCount":
		if self.present != nil {
			if _, ok := self.present["executorIdsCount"]; ok {
				return self.ExecutorIdsCount, nil
			}
		}
		return nil, fmt.Errorf("Field ExecutorIdsCount no set on ExecutorIdsCount %+v", self)

	case "executorIdsList", "ExecutorIdsList":
		if self.present != nil {
			if _, ok := self.present["executorIdsList"]; ok {
				return self.ExecutorIdsList, nil
			}
		}
		return nil, fmt.Errorf("Field ExecutorIdsList no set on ExecutorIdsList %+v", self)

	case "frameworkId", "FrameworkId":
		if self.present != nil {
			if _, ok := self.present["frameworkId"]; ok {
				return self.FrameworkId, nil
			}
		}
		return nil, fmt.Errorf("Field FrameworkId no set on FrameworkId %+v", self)

	case "frameworkIdOrBuilder", "FrameworkIdOrBuilder":
		if self.present != nil {
			if _, ok := self.present["frameworkIdOrBuilder"]; ok {
				return self.FrameworkIdOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field FrameworkIdOrBuilder no set on FrameworkIdOrBuilder %+v", self)

	case "hostname", "Hostname":
		if self.present != nil {
			if _, ok := self.present["hostname"]; ok {
				return self.Hostname, nil
			}
		}
		return nil, fmt.Errorf("Field Hostname no set on Hostname %+v", self)

	case "hostnameBytes", "HostnameBytes":
		if self.present != nil {
			if _, ok := self.present["hostnameBytes"]; ok {
				return self.HostnameBytes, nil
			}
		}
		return nil, fmt.Errorf("Field HostnameBytes no set on HostnameBytes %+v", self)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "idOrBuilder", "IdOrBuilder":
		if self.present != nil {
			if _, ok := self.present["idOrBuilder"]; ok {
				return self.IdOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field IdOrBuilder no set on IdOrBuilder %+v", self)

	case "initializationErrorString", "InitializationErrorString":
		if self.present != nil {
			if _, ok := self.present["initializationErrorString"]; ok {
				return self.InitializationErrorString, nil
			}
		}
		return nil, fmt.Errorf("Field InitializationErrorString no set on InitializationErrorString %+v", self)

	case "initialized", "Initialized":
		if self.present != nil {
			if _, ok := self.present["initialized"]; ok {
				return self.Initialized, nil
			}
		}
		return nil, fmt.Errorf("Field Initialized no set on Initialized %+v", self)

	case "resourcesCount", "ResourcesCount":
		if self.present != nil {
			if _, ok := self.present["resourcesCount"]; ok {
				return self.ResourcesCount, nil
			}
		}
		return nil, fmt.Errorf("Field ResourcesCount no set on ResourcesCount %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "slaveId", "SlaveId":
		if self.present != nil {
			if _, ok := self.present["slaveId"]; ok {
				return self.SlaveId, nil
			}
		}
		return nil, fmt.Errorf("Field SlaveId no set on SlaveId %+v", self)

	case "slaveIdOrBuilder", "SlaveIdOrBuilder":
		if self.present != nil {
			if _, ok := self.present["slaveIdOrBuilder"]; ok {
				return self.SlaveIdOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field SlaveIdOrBuilder no set on SlaveIdOrBuilder %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *Offer) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Offer", name)

	case "attributesCount", "AttributesCount":
		self.present["attributesCount"] = false

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "executorIdsCount", "ExecutorIdsCount":
		self.present["executorIdsCount"] = false

	case "executorIdsList", "ExecutorIdsList":
		self.present["executorIdsList"] = false

	case "frameworkId", "FrameworkId":
		self.present["frameworkId"] = false

	case "frameworkIdOrBuilder", "FrameworkIdOrBuilder":
		self.present["frameworkIdOrBuilder"] = false

	case "hostname", "Hostname":
		self.present["hostname"] = false

	case "hostnameBytes", "HostnameBytes":
		self.present["hostnameBytes"] = false

	case "id", "Id":
		self.present["id"] = false

	case "idOrBuilder", "IdOrBuilder":
		self.present["idOrBuilder"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "resourcesCount", "ResourcesCount":
		self.present["resourcesCount"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "slaveId", "SlaveId":
		self.present["slaveId"] = false

	case "slaveIdOrBuilder", "SlaveIdOrBuilder":
		self.present["slaveIdOrBuilder"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *Offer) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type OfferList []*Offer

func (self *OfferList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*OfferList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A OfferList cannot copy the values from %#v", other)
}

func (list *OfferList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *OfferList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *OfferList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
