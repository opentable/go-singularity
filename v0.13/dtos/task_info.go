package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type TaskInfo struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	Command *CommandInfo `json:"command"`

	CommandOrBuilder *CommandInfoOrBuilder `json:"commandOrBuilder"`

	Container *ContainerInfo `json:"container"`

	ContainerOrBuilder *ContainerInfoOrBuilder `json:"containerOrBuilder"`

	Data *ByteString `json:"data"`

	DefaultInstanceForType *TaskInfo `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Discovery *DiscoveryInfo `json:"discovery"`

	DiscoveryOrBuilder *DiscoveryInfoOrBuilder `json:"discoveryOrBuilder"`

	Executor *ExecutorInfo `json:"executor"`

	ExecutorOrBuilder *ExecutorInfoOrBuilder `json:"executorOrBuilder"`

	HealthCheck *HealthCheck `json:"healthCheck"`

	HealthCheckOrBuilder *HealthCheckOrBuilder `json:"healthCheckOrBuilder"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	Labels *Labels `json:"labels"`

	LabelsOrBuilder *LabelsOrBuilder `json:"labelsOrBuilder"`

	Name string `json:"name,omitempty"`

	NameBytes *ByteString `json:"nameBytes"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$TaskInfo> `json:"parserForType"`

	ResourcesCount int32 `json:"resourcesCount"`

	// ResourcesList *List[Resource] `json:"resourcesList"`

	// ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder] `json:"resourcesOrBuilderList"`

	SerializedSize int32 `json:"serializedSize"`

	SlaveId *SlaveID `json:"slaveId"`

	SlaveIdOrBuilder *SlaveIDOrBuilder `json:"slaveIdOrBuilder"`

	TaskId *TaskID `json:"taskId"`

	TaskIdOrBuilder *TaskIDOrBuilder `json:"taskIdOrBuilder"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *TaskInfo) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *TaskInfo) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*TaskInfo); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A TaskInfo cannot copy the values from %#v", other)
}

func (self *TaskInfo) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *TaskInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *TaskInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *TaskInfo) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *TaskInfo) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on TaskInfo", name)

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
		v, ok := value.(*TaskInfo)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *TaskInfo", value, value)
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

	case "executor", "Executor":
		v, ok := value.(*ExecutorInfo)
		if ok {
			self.Executor = v
			self.present["executor"] = true
			return nil
		} else {
			return fmt.Errorf("Field executor/Executor: value %v(%T) couldn't be cast to type *ExecutorInfo", value, value)
		}

	case "executorOrBuilder", "ExecutorOrBuilder":
		v, ok := value.(*ExecutorInfoOrBuilder)
		if ok {
			self.ExecutorOrBuilder = v
			self.present["executorOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field executorOrBuilder/ExecutorOrBuilder: value %v(%T) couldn't be cast to type *ExecutorInfoOrBuilder", value, value)
		}

	case "healthCheck", "HealthCheck":
		v, ok := value.(*HealthCheck)
		if ok {
			self.HealthCheck = v
			self.present["healthCheck"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthCheck/HealthCheck: value %v(%T) couldn't be cast to type *HealthCheck", value, value)
		}

	case "healthCheckOrBuilder", "HealthCheckOrBuilder":
		v, ok := value.(*HealthCheckOrBuilder)
		if ok {
			self.HealthCheckOrBuilder = v
			self.present["healthCheckOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthCheckOrBuilder/HealthCheckOrBuilder: value %v(%T) couldn't be cast to type *HealthCheckOrBuilder", value, value)
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

	case "labels", "Labels":
		v, ok := value.(*Labels)
		if ok {
			self.Labels = v
			self.present["labels"] = true
			return nil
		} else {
			return fmt.Errorf("Field labels/Labels: value %v(%T) couldn't be cast to type *Labels", value, value)
		}

	case "labelsOrBuilder", "LabelsOrBuilder":
		v, ok := value.(*LabelsOrBuilder)
		if ok {
			self.LabelsOrBuilder = v
			self.present["labelsOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field labelsOrBuilder/LabelsOrBuilder: value %v(%T) couldn't be cast to type *LabelsOrBuilder", value, value)
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

	case "taskId", "TaskId":
		v, ok := value.(*TaskID)
		if ok {
			self.TaskId = v
			self.present["taskId"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *TaskID", value, value)
		}

	case "taskIdOrBuilder", "TaskIdOrBuilder":
		v, ok := value.(*TaskIDOrBuilder)
		if ok {
			self.TaskIdOrBuilder = v
			self.present["taskIdOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskIdOrBuilder/TaskIdOrBuilder: value %v(%T) couldn't be cast to type *TaskIDOrBuilder", value, value)
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

func (self *TaskInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on TaskInfo", name)

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

	case "executor", "Executor":
		if self.present != nil {
			if _, ok := self.present["executor"]; ok {
				return self.Executor, nil
			}
		}
		return nil, fmt.Errorf("Field Executor no set on Executor %+v", self)

	case "executorOrBuilder", "ExecutorOrBuilder":
		if self.present != nil {
			if _, ok := self.present["executorOrBuilder"]; ok {
				return self.ExecutorOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field ExecutorOrBuilder no set on ExecutorOrBuilder %+v", self)

	case "healthCheck", "HealthCheck":
		if self.present != nil {
			if _, ok := self.present["healthCheck"]; ok {
				return self.HealthCheck, nil
			}
		}
		return nil, fmt.Errorf("Field HealthCheck no set on HealthCheck %+v", self)

	case "healthCheckOrBuilder", "HealthCheckOrBuilder":
		if self.present != nil {
			if _, ok := self.present["healthCheckOrBuilder"]; ok {
				return self.HealthCheckOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field HealthCheckOrBuilder no set on HealthCheckOrBuilder %+v", self)

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

	case "labels", "Labels":
		if self.present != nil {
			if _, ok := self.present["labels"]; ok {
				return self.Labels, nil
			}
		}
		return nil, fmt.Errorf("Field Labels no set on Labels %+v", self)

	case "labelsOrBuilder", "LabelsOrBuilder":
		if self.present != nil {
			if _, ok := self.present["labelsOrBuilder"]; ok {
				return self.LabelsOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field LabelsOrBuilder no set on LabelsOrBuilder %+v", self)

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

	case "taskId", "TaskId":
		if self.present != nil {
			if _, ok := self.present["taskId"]; ok {
				return self.TaskId, nil
			}
		}
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "taskIdOrBuilder", "TaskIdOrBuilder":
		if self.present != nil {
			if _, ok := self.present["taskIdOrBuilder"]; ok {
				return self.TaskIdOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field TaskIdOrBuilder no set on TaskIdOrBuilder %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *TaskInfo) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on TaskInfo", name)

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

	case "executor", "Executor":
		self.present["executor"] = false

	case "executorOrBuilder", "ExecutorOrBuilder":
		self.present["executorOrBuilder"] = false

	case "healthCheck", "HealthCheck":
		self.present["healthCheck"] = false

	case "healthCheckOrBuilder", "HealthCheckOrBuilder":
		self.present["healthCheckOrBuilder"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "labels", "Labels":
		self.present["labels"] = false

	case "labelsOrBuilder", "LabelsOrBuilder":
		self.present["labelsOrBuilder"] = false

	case "name", "Name":
		self.present["name"] = false

	case "nameBytes", "NameBytes":
		self.present["nameBytes"] = false

	case "resourcesCount", "ResourcesCount":
		self.present["resourcesCount"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "slaveId", "SlaveId":
		self.present["slaveId"] = false

	case "slaveIdOrBuilder", "SlaveIdOrBuilder":
		self.present["slaveIdOrBuilder"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	case "taskIdOrBuilder", "TaskIdOrBuilder":
		self.present["taskIdOrBuilder"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *TaskInfo) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type TaskInfoList []*TaskInfo

func (self *TaskInfoList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*TaskInfoList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A TaskInfoList cannot copy the values from %#v", other)
}

func (list *TaskInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *TaskInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *TaskInfoList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
