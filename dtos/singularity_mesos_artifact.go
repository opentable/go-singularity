package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityMesosArtifact struct {
	present map[string]bool

	Uri string `json:"uri,omitempty"`

	Cache bool `json:"cache"`

	Executable bool `json:"executable"`

	Extract bool `json:"extract"`
}

func (self *SingularityMesosArtifact) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityMesosArtifact) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityMesosArtifact); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityMesosArtifact cannot copy the values from %#v", other)
}

func (self *SingularityMesosArtifact) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityMesosArtifact) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityMesosArtifact) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityMesosArtifact) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityMesosArtifact) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMesosArtifact", name)

	case "uri", "Uri":
		v, ok := value.(string)
		if ok {
			self.Uri = v
			self.present["uri"] = true
			return nil
		} else {
			return fmt.Errorf("Field uri/Uri: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "cache", "Cache":
		v, ok := value.(bool)
		if ok {
			self.Cache = v
			self.present["cache"] = true
			return nil
		} else {
			return fmt.Errorf("Field cache/Cache: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "executable", "Executable":
		v, ok := value.(bool)
		if ok {
			self.Executable = v
			self.present["executable"] = true
			return nil
		} else {
			return fmt.Errorf("Field executable/Executable: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "extract", "Extract":
		v, ok := value.(bool)
		if ok {
			self.Extract = v
			self.present["extract"] = true
			return nil
		} else {
			return fmt.Errorf("Field extract/Extract: value %v(%T) couldn't be cast to type bool", value, value)
		}

	}
}

func (self *SingularityMesosArtifact) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityMesosArtifact", name)

	case "uri", "Uri":
		if self.present != nil {
			if _, ok := self.present["uri"]; ok {
				return self.Uri, nil
			}
		}
		return nil, fmt.Errorf("Field Uri no set on Uri %+v", self)

	case "cache", "Cache":
		if self.present != nil {
			if _, ok := self.present["cache"]; ok {
				return self.Cache, nil
			}
		}
		return nil, fmt.Errorf("Field Cache no set on Cache %+v", self)

	case "executable", "Executable":
		if self.present != nil {
			if _, ok := self.present["executable"]; ok {
				return self.Executable, nil
			}
		}
		return nil, fmt.Errorf("Field Executable no set on Executable %+v", self)

	case "extract", "Extract":
		if self.present != nil {
			if _, ok := self.present["extract"]; ok {
				return self.Extract, nil
			}
		}
		return nil, fmt.Errorf("Field Extract no set on Extract %+v", self)

	}
}

func (self *SingularityMesosArtifact) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMesosArtifact", name)

	case "uri", "Uri":
		self.present["uri"] = false

	case "cache", "Cache":
		self.present["cache"] = false

	case "executable", "Executable":
		self.present["executable"] = false

	case "extract", "Extract":
		self.present["extract"] = false

	}

	return nil
}

func (self *SingularityMesosArtifact) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityMesosArtifactList []*SingularityMesosArtifact

func (self *SingularityMesosArtifactList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityMesosArtifactList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityMesosArtifactList cannot copy the values from %#v", other)
}

func (list *SingularityMesosArtifactList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityMesosArtifactList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityMesosArtifactList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
