package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ExecutorInfo struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	Command *CommandInfo `json:"command"`

	CommandOrBuilder *CommandInfoOrBuilder `json:"commandOrBuilder"`

	Container *ContainerInfo `json:"container"`

	ContainerOrBuilder *ContainerInfoOrBuilder `json:"containerOrBuilder"`

	Data *ByteString `json:"data"`

	DefaultInstanceForType *ExecutorInfo `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Discovery *DiscoveryInfo `json:"discovery"`

	DiscoveryOrBuilder *DiscoveryInfoOrBuilder `json:"discoveryOrBuilder"`

	ExecutorId *ExecutorID `json:"executorId"`

	ExecutorIdOrBuilder *ExecutorIDOrBuilder `json:"executorIdOrBuilder"`

	FrameworkId *FrameworkID `json:"frameworkId"`

	FrameworkIdOrBuilder *FrameworkIDOrBuilder `json:"frameworkIdOrBuilder"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	Name string `json:"name,omitempty"`

	NameBytes *ByteString `json:"nameBytes"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$ExecutorInfo> `json:"parserForType"`

	ResourcesCount int32 `json:"resourcesCount"`

	// ResourcesList *List[Resource] `json:"resourcesList"`

	// ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder] `json:"resourcesOrBuilderList"`

	SerializedSize int32 `json:"serializedSize"`

	Source string `json:"source,omitempty"`

	SourceBytes *ByteString `json:"sourceBytes"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *ExecutorInfo) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *ExecutorInfo) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ExecutorInfo); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ExecutorInfo cannot copy the values from %#v", other)
}

func (self *ExecutorInfo) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *ExecutorInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ExecutorInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ExecutorInfo) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *ExecutorInfo) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExecutorInfo", name)

	case "command", "Command":
		v, ok := value.(*CommandInfo)
		if ok {
			self.Command = v
			self.present["command"] = true
			return nil
		} else {
			return fmt.Errorf("Field command/Command: value %v(%T) couldn't be cast to type *CommandInfo", value, value)
		}

	case "commandOrBuilder", "CommandOrBuilder":
		v, ok := value.(*CommandInfoOrBuilder)
		if ok {
			self.CommandOrBuilder = v
			self.present["commandOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field commandOrBuilder/CommandOrBuilder: value %v(%T) couldn't be cast to type *CommandInfoOrBuilder", value, value)
		}

	case "container", "Container":
		v, ok := value.(*ContainerInfo)
		if ok {
			self.Container = v
			self.present["container"] = true
			return nil
		} else {
			return fmt.Errorf("Field container/Container: value %v(%T) couldn't be cast to type *ContainerInfo", value, value)
		}

	case "containerOrBuilder", "ContainerOrBuilder":
		v, ok := value.(*ContainerInfoOrBuilder)
		if ok {
			self.ContainerOrBuilder = v
			self.present["containerOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field containerOrBuilder/ContainerOrBuilder: value %v(%T) couldn't be cast to type *ContainerInfoOrBuilder", value, value)
		}

	case "data", "Data":
		v, ok := value.(*ByteString)
		if ok {
			self.Data = v
			self.present["data"] = true
			return nil
		} else {
			return fmt.Errorf("Field data/Data: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*ExecutorInfo)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *ExecutorInfo", value, value)
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

	case "discovery", "Discovery":
		v, ok := value.(*DiscoveryInfo)
		if ok {
			self.Discovery = v
			self.present["discovery"] = true
			return nil
		} else {
			return fmt.Errorf("Field discovery/Discovery: value %v(%T) couldn't be cast to type *DiscoveryInfo", value, value)
		}

	case "discoveryOrBuilder", "DiscoveryOrBuilder":
		v, ok := value.(*DiscoveryInfoOrBuilder)
		if ok {
			self.DiscoveryOrBuilder = v
			self.present["discoveryOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field discoveryOrBuilder/DiscoveryOrBuilder: value %v(%T) couldn't be cast to type *DiscoveryInfoOrBuilder", value, value)
		}

	case "executorId", "ExecutorId":
		v, ok := value.(*ExecutorID)
		if ok {
			self.ExecutorId = v
			self.present["executorId"] = true
			return nil
		} else {
			return fmt.Errorf("Field executorId/ExecutorId: value %v(%T) couldn't be cast to type *ExecutorID", value, value)
		}

	case "executorIdOrBuilder", "ExecutorIdOrBuilder":
		v, ok := value.(*ExecutorIDOrBuilder)
		if ok {
			self.ExecutorIdOrBuilder = v
			self.present["executorIdOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field executorIdOrBuilder/ExecutorIdOrBuilder: value %v(%T) couldn't be cast to type *ExecutorIDOrBuilder", value, value)
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

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = v
			self.present["name"] = true
			return nil
		} else {
			return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "nameBytes", "NameBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.NameBytes = v
			self.present["nameBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field nameBytes/NameBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	case "source", "Source":
		v, ok := value.(string)
		if ok {
			self.Source = v
			self.present["source"] = true
			return nil
		} else {
			return fmt.Errorf("Field source/Source: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "sourceBytes", "SourceBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.SourceBytes = v
			self.present["sourceBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field sourceBytes/SourceBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

func (self *ExecutorInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ExecutorInfo", name)

	case "command", "Command":
		if self.present != nil {
			if _, ok := self.present["command"]; ok {
				return self.Command, nil
			}
		}
		return nil, fmt.Errorf("Field Command no set on Command %+v", self)

	case "commandOrBuilder", "CommandOrBuilder":
		if self.present != nil {
			if _, ok := self.present["commandOrBuilder"]; ok {
				return self.CommandOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field CommandOrBuilder no set on CommandOrBuilder %+v", self)

	case "container", "Container":
		if self.present != nil {
			if _, ok := self.present["container"]; ok {
				return self.Container, nil
			}
		}
		return nil, fmt.Errorf("Field Container no set on Container %+v", self)

	case "containerOrBuilder", "ContainerOrBuilder":
		if self.present != nil {
			if _, ok := self.present["containerOrBuilder"]; ok {
				return self.ContainerOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field ContainerOrBuilder no set on ContainerOrBuilder %+v", self)

	case "data", "Data":
		if self.present != nil {
			if _, ok := self.present["data"]; ok {
				return self.Data, nil
			}
		}
		return nil, fmt.Errorf("Field Data no set on Data %+v", self)

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

	case "discovery", "Discovery":
		if self.present != nil {
			if _, ok := self.present["discovery"]; ok {
				return self.Discovery, nil
			}
		}
		return nil, fmt.Errorf("Field Discovery no set on Discovery %+v", self)

	case "discoveryOrBuilder", "DiscoveryOrBuilder":
		if self.present != nil {
			if _, ok := self.present["discoveryOrBuilder"]; ok {
				return self.DiscoveryOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field DiscoveryOrBuilder no set on DiscoveryOrBuilder %+v", self)

	case "executorId", "ExecutorId":
		if self.present != nil {
			if _, ok := self.present["executorId"]; ok {
				return self.ExecutorId, nil
			}
		}
		return nil, fmt.Errorf("Field ExecutorId no set on ExecutorId %+v", self)

	case "executorIdOrBuilder", "ExecutorIdOrBuilder":
		if self.present != nil {
			if _, ok := self.present["executorIdOrBuilder"]; ok {
				return self.ExecutorIdOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field ExecutorIdOrBuilder no set on ExecutorIdOrBuilder %+v", self)

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

	case "name", "Name":
		if self.present != nil {
			if _, ok := self.present["name"]; ok {
				return self.Name, nil
			}
		}
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "nameBytes", "NameBytes":
		if self.present != nil {
			if _, ok := self.present["nameBytes"]; ok {
				return self.NameBytes, nil
			}
		}
		return nil, fmt.Errorf("Field NameBytes no set on NameBytes %+v", self)

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

	case "source", "Source":
		if self.present != nil {
			if _, ok := self.present["source"]; ok {
				return self.Source, nil
			}
		}
		return nil, fmt.Errorf("Field Source no set on Source %+v", self)

	case "sourceBytes", "SourceBytes":
		if self.present != nil {
			if _, ok := self.present["sourceBytes"]; ok {
				return self.SourceBytes, nil
			}
		}
		return nil, fmt.Errorf("Field SourceBytes no set on SourceBytes %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *ExecutorInfo) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExecutorInfo", name)

	case "command", "Command":
		self.present["command"] = false

	case "commandOrBuilder", "CommandOrBuilder":
		self.present["commandOrBuilder"] = false

	case "container", "Container":
		self.present["container"] = false

	case "containerOrBuilder", "ContainerOrBuilder":
		self.present["containerOrBuilder"] = false

	case "data", "Data":
		self.present["data"] = false

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "discovery", "Discovery":
		self.present["discovery"] = false

	case "discoveryOrBuilder", "DiscoveryOrBuilder":
		self.present["discoveryOrBuilder"] = false

	case "executorId", "ExecutorId":
		self.present["executorId"] = false

	case "executorIdOrBuilder", "ExecutorIdOrBuilder":
		self.present["executorIdOrBuilder"] = false

	case "frameworkId", "FrameworkId":
		self.present["frameworkId"] = false

	case "frameworkIdOrBuilder", "FrameworkIdOrBuilder":
		self.present["frameworkIdOrBuilder"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "name", "Name":
		self.present["name"] = false

	case "nameBytes", "NameBytes":
		self.present["nameBytes"] = false

	case "resourcesCount", "ResourcesCount":
		self.present["resourcesCount"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "source", "Source":
		self.present["source"] = false

	case "sourceBytes", "SourceBytes":
		self.present["sourceBytes"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *ExecutorInfo) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ExecutorInfoList []*ExecutorInfo

func (self *ExecutorInfoList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ExecutorInfoList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ExecutorInfoList cannot copy the values from %#v", other)
}

func (list *ExecutorInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *ExecutorInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ExecutorInfoList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
