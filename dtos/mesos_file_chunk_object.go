package dtos

import (
	"fmt"
	"io"
)

type MesosFileChunkObject struct {
	present    map[string]bool
	Data       string `json:"data,omitempty"`
	NextOffset int64  `json:"nextOffset"`
	Offset     int64  `json:"offset"`
}

func (self *MesosFileChunkObject) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *MesosFileChunkObject) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *MesosFileChunkObject) FormatText() string {
	return FormatText(self)
}

func (self *MesosFileChunkObject) FormatJSON() string {
	return FormatJSON(self)
}

func (self *MesosFileChunkObject) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *MesosFileChunkObject) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosFileChunkObject", name)

	case "data", "Data":
		v, ok := value.(string)
		if ok {
			self.Data = v
			self.present["data"] = true
			return nil
		} else {
			return fmt.Errorf("Field data/Data: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "nextOffset", "NextOffset":
		v, ok := value.(int64)
		if ok {
			self.NextOffset = v
			self.present["nextOffset"] = true
			return nil
		} else {
			return fmt.Errorf("Field nextOffset/NextOffset: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "offset", "Offset":
		v, ok := value.(int64)
		if ok {
			self.Offset = v
			self.present["offset"] = true
			return nil
		} else {
			return fmt.Errorf("Field offset/Offset: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *MesosFileChunkObject) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MesosFileChunkObject", name)

	case "data", "Data":
		if self.present != nil {
			if _, ok := self.present["data"]; ok {
				return self.Data, nil
			}
		}
		return nil, fmt.Errorf("Field Data no set on Data %+v", self)

	case "nextOffset", "NextOffset":
		if self.present != nil {
			if _, ok := self.present["nextOffset"]; ok {
				return self.NextOffset, nil
			}
		}
		return nil, fmt.Errorf("Field NextOffset no set on NextOffset %+v", self)

	case "offset", "Offset":
		if self.present != nil {
			if _, ok := self.present["offset"]; ok {
				return self.Offset, nil
			}
		}
		return nil, fmt.Errorf("Field Offset no set on Offset %+v", self)

	}
}

func (self *MesosFileChunkObject) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosFileChunkObject", name)

	case "data", "Data":
		self.present["data"] = false

	case "nextOffset", "NextOffset":
		self.present["nextOffset"] = false

	case "offset", "Offset":
		self.present["offset"] = false

	}

	return nil
}

func (self *MesosFileChunkObject) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type MesosFileChunkObjectList []*MesosFileChunkObject

func (list *MesosFileChunkObjectList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *MesosFileChunkObjectList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MesosFileChunkObjectList) FormatJSON() string {
	return FormatJSON(list)
}
