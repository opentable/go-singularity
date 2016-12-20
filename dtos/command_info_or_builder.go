package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type CommandInfoOrBuilder struct {
	present map[string]bool

	ArgumentsCount int32 `json:"argumentsCount"`

	ArgumentsList swaggering.StringList `json:"argumentsList"`

	Container *ContainerInfo `json:"container"`

	ContainerOrBuilder *ContainerInfoOrBuilder `json:"containerOrBuilder"`

	Environment *Environment `json:"environment"`

	EnvironmentOrBuilder *EnvironmentOrBuilder `json:"environmentOrBuilder"`

	Shell bool `json:"shell"`

	UrisCount int32 `json:"urisCount"`

	// UrisList *List[URI] `json:"urisList"`

	// UrisOrBuilderList *List[? extends org.apache.mesos.Protos$CommandInfo$URIOrBuilder] `json:"urisOrBuilderList"`

	User string `json:"user,omitempty"`

	UserBytes *ByteString `json:"userBytes"`

	Value string `json:"value,omitempty"`

	ValueBytes *ByteString `json:"valueBytes"`
}

func (self *CommandInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *CommandInfoOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*CommandInfoOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A CommandInfoOrBuilder cannot copy the values from %#v", other)
}

func (self *CommandInfoOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *CommandInfoOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *CommandInfoOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *CommandInfoOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *CommandInfoOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on CommandInfoOrBuilder", name)

	case "argumentsCount", "ArgumentsCount":
		v, ok := value.(int32)
		if ok {
			self.ArgumentsCount = v
			self.present["argumentsCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field argumentsCount/ArgumentsCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "argumentsList", "ArgumentsList":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.ArgumentsList = v
			self.present["argumentsList"] = true
			return nil
		} else {
			return fmt.Errorf("Field argumentsList/ArgumentsList: value %v(%T) couldn't be cast to type StringList", value, value)
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

	case "environment", "Environment":
		v, ok := value.(*Environment)
		if ok {
			self.Environment = v
			self.present["environment"] = true
			return nil
		} else {
			return fmt.Errorf("Field environment/Environment: value %v(%T) couldn't be cast to type *Environment", value, value)
		}

	case "environmentOrBuilder", "EnvironmentOrBuilder":
		v, ok := value.(*EnvironmentOrBuilder)
		if ok {
			self.EnvironmentOrBuilder = v
			self.present["environmentOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field environmentOrBuilder/EnvironmentOrBuilder: value %v(%T) couldn't be cast to type *EnvironmentOrBuilder", value, value)
		}

	case "shell", "Shell":
		v, ok := value.(bool)
		if ok {
			self.Shell = v
			self.present["shell"] = true
			return nil
		} else {
			return fmt.Errorf("Field shell/Shell: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "urisCount", "UrisCount":
		v, ok := value.(int32)
		if ok {
			self.UrisCount = v
			self.present["urisCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field urisCount/UrisCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "userBytes", "UserBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.UserBytes = v
			self.present["userBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field userBytes/UserBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "value", "Value":
		v, ok := value.(string)
		if ok {
			self.Value = v
			self.present["value"] = true
			return nil
		} else {
			return fmt.Errorf("Field value/Value: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "valueBytes", "ValueBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.ValueBytes = v
			self.present["valueBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field valueBytes/ValueBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	}
}

func (self *CommandInfoOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on CommandInfoOrBuilder", name)

	case "argumentsCount", "ArgumentsCount":
		if self.present != nil {
			if _, ok := self.present["argumentsCount"]; ok {
				return self.ArgumentsCount, nil
			}
		}
		return nil, fmt.Errorf("Field ArgumentsCount no set on ArgumentsCount %+v", self)

	case "argumentsList", "ArgumentsList":
		if self.present != nil {
			if _, ok := self.present["argumentsList"]; ok {
				return self.ArgumentsList, nil
			}
		}
		return nil, fmt.Errorf("Field ArgumentsList no set on ArgumentsList %+v", self)

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

	case "environment", "Environment":
		if self.present != nil {
			if _, ok := self.present["environment"]; ok {
				return self.Environment, nil
			}
		}
		return nil, fmt.Errorf("Field Environment no set on Environment %+v", self)

	case "environmentOrBuilder", "EnvironmentOrBuilder":
		if self.present != nil {
			if _, ok := self.present["environmentOrBuilder"]; ok {
				return self.EnvironmentOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field EnvironmentOrBuilder no set on EnvironmentOrBuilder %+v", self)

	case "shell", "Shell":
		if self.present != nil {
			if _, ok := self.present["shell"]; ok {
				return self.Shell, nil
			}
		}
		return nil, fmt.Errorf("Field Shell no set on Shell %+v", self)

	case "urisCount", "UrisCount":
		if self.present != nil {
			if _, ok := self.present["urisCount"]; ok {
				return self.UrisCount, nil
			}
		}
		return nil, fmt.Errorf("Field UrisCount no set on UrisCount %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "userBytes", "UserBytes":
		if self.present != nil {
			if _, ok := self.present["userBytes"]; ok {
				return self.UserBytes, nil
			}
		}
		return nil, fmt.Errorf("Field UserBytes no set on UserBytes %+v", self)

	case "value", "Value":
		if self.present != nil {
			if _, ok := self.present["value"]; ok {
				return self.Value, nil
			}
		}
		return nil, fmt.Errorf("Field Value no set on Value %+v", self)

	case "valueBytes", "ValueBytes":
		if self.present != nil {
			if _, ok := self.present["valueBytes"]; ok {
				return self.ValueBytes, nil
			}
		}
		return nil, fmt.Errorf("Field ValueBytes no set on ValueBytes %+v", self)

	}
}

func (self *CommandInfoOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on CommandInfoOrBuilder", name)

	case "argumentsCount", "ArgumentsCount":
		self.present["argumentsCount"] = false

	case "argumentsList", "ArgumentsList":
		self.present["argumentsList"] = false

	case "container", "Container":
		self.present["container"] = false

	case "containerOrBuilder", "ContainerOrBuilder":
		self.present["containerOrBuilder"] = false

	case "environment", "Environment":
		self.present["environment"] = false

	case "environmentOrBuilder", "EnvironmentOrBuilder":
		self.present["environmentOrBuilder"] = false

	case "shell", "Shell":
		self.present["shell"] = false

	case "urisCount", "UrisCount":
		self.present["urisCount"] = false

	case "user", "User":
		self.present["user"] = false

	case "userBytes", "UserBytes":
		self.present["userBytes"] = false

	case "value", "Value":
		self.present["value"] = false

	case "valueBytes", "ValueBytes":
		self.present["valueBytes"] = false

	}

	return nil
}

func (self *CommandInfoOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type CommandInfoOrBuilderList []*CommandInfoOrBuilder

func (self *CommandInfoOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*CommandInfoOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A CommandInfoOrBuilderList cannot copy the values from %#v", other)
}

func (list *CommandInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *CommandInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *CommandInfoOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
