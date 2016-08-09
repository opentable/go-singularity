package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type AppcOrBuilder struct {
	present map[string]bool

	Id string `json:"id,omitempty"`

	IdBytes *ByteString `json:"idBytes"`

	Labels *Labels `json:"labels"`

	LabelsOrBuilder *LabelsOrBuilder `json:"labelsOrBuilder"`

	Name string `json:"name,omitempty"`

	NameBytes *ByteString `json:"nameBytes"`
}

func (self *AppcOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *AppcOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*AppcOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A AppcOrBuilder cannot absorb the values from %v", other)
}

func (self *AppcOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *AppcOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *AppcOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *AppcOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *AppcOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on AppcOrBuilder", name)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "idBytes", "IdBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.IdBytes = v
			self.present["idBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field idBytes/IdBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	}
}

func (self *AppcOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on AppcOrBuilder", name)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "idBytes", "IdBytes":
		if self.present != nil {
			if _, ok := self.present["idBytes"]; ok {
				return self.IdBytes, nil
			}
		}
		return nil, fmt.Errorf("Field IdBytes no set on IdBytes %+v", self)

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

	}
}

func (self *AppcOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on AppcOrBuilder", name)

	case "id", "Id":
		self.present["id"] = false

	case "idBytes", "IdBytes":
		self.present["idBytes"] = false

	case "labels", "Labels":
		self.present["labels"] = false

	case "labelsOrBuilder", "LabelsOrBuilder":
		self.present["labelsOrBuilder"] = false

	case "name", "Name":
		self.present["name"] = false

	case "nameBytes", "NameBytes":
		self.present["nameBytes"] = false

	}

	return nil
}

func (self *AppcOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type AppcOrBuilderList []*AppcOrBuilder

func (self *AppcOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*AppcOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A AppcOrBuilder cannot absorb the values from %v", other)
}

func (list *AppcOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *AppcOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *AppcOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
