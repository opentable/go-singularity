package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ExecutorInfoOrBuilder struct {
	present map[string]bool

	Command *CommandInfo `json:"command"`

	CommandOrBuilder *CommandInfoOrBuilder `json:"commandOrBuilder"`

	Container *ContainerInfo `json:"container"`

	ContainerOrBuilder *ContainerInfoOrBuilder `json:"containerOrBuilder"`

	Data *ByteString `json:"data"`

	Discovery *DiscoveryInfo `json:"discovery"`

	DiscoveryOrBuilder *DiscoveryInfoOrBuilder `json:"discoveryOrBuilder"`

	ExecutorId *ExecutorID `json:"executorId"`

	ExecutorIdOrBuilder *ExecutorIDOrBuilder `json:"executorIdOrBuilder"`

	FrameworkId *FrameworkID `json:"frameworkId"`

	FrameworkIdOrBuilder *FrameworkIDOrBuilder `json:"frameworkIdOrBuilder"`

	Name string `json:"name,omitempty"`

	NameBytes *ByteString `json:"nameBytes"`

	ResourcesCount int32 `json:"resourcesCount"`

	// ResourcesList *List[Resource] `json:"resourcesList"`

	// ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder] `json:"resourcesOrBuilderList"`

	Source string `json:"source,omitempty"`

	SourceBytes *ByteString `json:"sourceBytes"`
}

func (self *ExecutorInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *ExecutorInfoOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ExecutorInfoOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ExecutorInfoOrBuilder cannot copy the values from %#v", other)
}

func (self *ExecutorInfoOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *ExecutorInfoOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ExecutorInfoOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ExecutorInfoOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *ExecutorInfoOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExecutorInfoOrBuilder", name)

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

	}
}

func (self *ExecutorInfoOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ExecutorInfoOrBuilder", name)

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

	}
}

func (self *ExecutorInfoOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExecutorInfoOrBuilder", name)

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

	case "name", "Name":
		self.present["name"] = false

	case "nameBytes", "NameBytes":
		self.present["nameBytes"] = false

	case "resourcesCount", "ResourcesCount":
		self.present["resourcesCount"] = false

	case "source", "Source":
		self.present["source"] = false

	case "sourceBytes", "SourceBytes":
		self.present["sourceBytes"] = false

	}

	return nil
}

func (self *ExecutorInfoOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ExecutorInfoOrBuilderList []*ExecutorInfoOrBuilder

func (self *ExecutorInfoOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ExecutorInfoOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ExecutorInfoOrBuilderList cannot copy the values from %#v", other)
}

func (list *ExecutorInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *ExecutorInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ExecutorInfoOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
