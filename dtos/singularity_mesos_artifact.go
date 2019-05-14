package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityMesosArtifact struct {
	Executable *bool   `json:"executable,omitempty"`
	Extract    *bool   `json:"extract,omitempty"`
	Uri        *string `json:"uri,omitempty"`
	Cache      *bool   `json:"cache,omitempty"`
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

func (self *SingularityMesosArtifact) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityMesosArtifact) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityMesosArtifact) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMesosArtifact", name)

	case "executable", "Executable":
		v, ok := value.(bool)
		if ok {
			self.Executable = &v
			return nil
		}
		return fmt.Errorf("Field executable/Executable: value %v(%T) couldn't be cast to type bool", value, value)

	case "extract", "Extract":
		v, ok := value.(bool)
		if ok {
			self.Extract = &v
			return nil
		}
		return fmt.Errorf("Field extract/Extract: value %v(%T) couldn't be cast to type bool", value, value)

	case "uri", "Uri":
		v, ok := value.(string)
		if ok {
			self.Uri = &v
			return nil
		}
		return fmt.Errorf("Field uri/Uri: value %v(%T) couldn't be cast to type string", value, value)

	case "cache", "Cache":
		v, ok := value.(bool)
		if ok {
			self.Cache = &v
			return nil
		}
		return fmt.Errorf("Field cache/Cache: value %v(%T) couldn't be cast to type bool", value, value)

	}
}

func (self *SingularityMesosArtifact) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityMesosArtifact", name)

	case "executable", "Executable":
		return *self.Executable, nil
		return nil, fmt.Errorf("Field Executable no set on Executable %+v", self)

	case "extract", "Extract":
		return *self.Extract, nil
		return nil, fmt.Errorf("Field Extract no set on Extract %+v", self)

	case "uri", "Uri":
		return *self.Uri, nil
		return nil, fmt.Errorf("Field Uri no set on Uri %+v", self)

	case "cache", "Cache":
		return *self.Cache, nil
		return nil, fmt.Errorf("Field Cache no set on Cache %+v", self)

	}
}

func (self *SingularityMesosArtifact) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMesosArtifact", name)

	case "executable", "Executable":
		self.Executable = nil

	case "extract", "Extract":
		self.Extract = nil

	case "uri", "Uri":
		self.Uri = nil

	case "cache", "Cache":
		self.Cache = nil

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
