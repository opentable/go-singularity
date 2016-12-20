package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type HTTPOrBuilder struct {
	present map[string]bool

	Path string `json:"path,omitempty"`

	PathBytes *ByteString `json:"pathBytes"`

	Port int32 `json:"port"`

	StatusesCount int32 `json:"statusesCount"`

	StatusesList []int32 `json:"statusesList"`
}

func (self *HTTPOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *HTTPOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HTTPOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HTTPOrBuilder cannot copy the values from %#v", other)
}

func (self *HTTPOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *HTTPOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *HTTPOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *HTTPOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *HTTPOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HTTPOrBuilder", name)

	case "path", "Path":
		v, ok := value.(string)
		if ok {
			self.Path = v
			self.present["path"] = true
			return nil
		} else {
			return fmt.Errorf("Field path/Path: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "pathBytes", "PathBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.PathBytes = v
			self.present["pathBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field pathBytes/PathBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "port", "Port":
		v, ok := value.(int32)
		if ok {
			self.Port = v
			self.present["port"] = true
			return nil
		} else {
			return fmt.Errorf("Field port/Port: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "statusesCount", "StatusesCount":
		v, ok := value.(int32)
		if ok {
			self.StatusesCount = v
			self.present["statusesCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field statusesCount/StatusesCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "statusesList", "StatusesList":
		v, ok := value.([]int32)
		if ok {
			self.StatusesList = v
			self.present["statusesList"] = true
			return nil
		} else {
			return fmt.Errorf("Field statusesList/StatusesList: value %v(%T) couldn't be cast to type []int32", value, value)
		}

	}
}

func (self *HTTPOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on HTTPOrBuilder", name)

	case "path", "Path":
		if self.present != nil {
			if _, ok := self.present["path"]; ok {
				return self.Path, nil
			}
		}
		return nil, fmt.Errorf("Field Path no set on Path %+v", self)

	case "pathBytes", "PathBytes":
		if self.present != nil {
			if _, ok := self.present["pathBytes"]; ok {
				return self.PathBytes, nil
			}
		}
		return nil, fmt.Errorf("Field PathBytes no set on PathBytes %+v", self)

	case "port", "Port":
		if self.present != nil {
			if _, ok := self.present["port"]; ok {
				return self.Port, nil
			}
		}
		return nil, fmt.Errorf("Field Port no set on Port %+v", self)

	case "statusesCount", "StatusesCount":
		if self.present != nil {
			if _, ok := self.present["statusesCount"]; ok {
				return self.StatusesCount, nil
			}
		}
		return nil, fmt.Errorf("Field StatusesCount no set on StatusesCount %+v", self)

	case "statusesList", "StatusesList":
		if self.present != nil {
			if _, ok := self.present["statusesList"]; ok {
				return self.StatusesList, nil
			}
		}
		return nil, fmt.Errorf("Field StatusesList no set on StatusesList %+v", self)

	}
}

func (self *HTTPOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HTTPOrBuilder", name)

	case "path", "Path":
		self.present["path"] = false

	case "pathBytes", "PathBytes":
		self.present["pathBytes"] = false

	case "port", "Port":
		self.present["port"] = false

	case "statusesCount", "StatusesCount":
		self.present["statusesCount"] = false

	case "statusesList", "StatusesList":
		self.present["statusesList"] = false

	}

	return nil
}

func (self *HTTPOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type HTTPOrBuilderList []*HTTPOrBuilder

func (self *HTTPOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HTTPOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HTTPOrBuilderList cannot copy the values from %#v", other)
}

func (list *HTTPOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *HTTPOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HTTPOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
