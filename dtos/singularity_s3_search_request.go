package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityS3SearchRequest struct {
	present map[string]bool

	ListOnly bool `json:"listOnly"`

	MaxPerPage int32 `json:"maxPerPage"`

	// Invalid field: ContinuationTokens *notfound.Map[string,ContinuationToken] `json:"continuationTokens"`

	RequestsAndDeploys map[string]swaggering.StringList `json:"requestsAndDeploys"`

	TaskIds swaggering.StringList `json:"taskIds"`

	Start int64 `json:"start"`

	End int64 `json:"end"`

	ExcludeMetadata bool `json:"excludeMetadata"`
}

func (self *SingularityS3SearchRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityS3SearchRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityS3SearchRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityS3SearchRequest cannot copy the values from %#v", other)
}

func (self *SingularityS3SearchRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityS3SearchRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityS3SearchRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityS3SearchRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityS3SearchRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityS3SearchRequest", name)

	case "listOnly", "ListOnly":
		v, ok := value.(bool)
		if ok {
			self.ListOnly = v
			self.present["listOnly"] = true
			return nil
		} else {
			return fmt.Errorf("Field listOnly/ListOnly: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "maxPerPage", "MaxPerPage":
		v, ok := value.(int32)
		if ok {
			self.MaxPerPage = v
			self.present["maxPerPage"] = true
			return nil
		} else {
			return fmt.Errorf("Field maxPerPage/MaxPerPage: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "requestsAndDeploys", "RequestsAndDeploys":
		v, ok := value.(map[string]swaggering.StringList)
		if ok {
			self.RequestsAndDeploys = v
			self.present["requestsAndDeploys"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestsAndDeploys/RequestsAndDeploys: value %v(%T) couldn't be cast to type map[string]swaggering.StringList", value, value)
		}

	case "taskIds", "TaskIds":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.TaskIds = v
			self.present["taskIds"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskIds/TaskIds: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)
		}

	case "start", "Start":
		v, ok := value.(int64)
		if ok {
			self.Start = v
			self.present["start"] = true
			return nil
		} else {
			return fmt.Errorf("Field start/Start: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "end", "End":
		v, ok := value.(int64)
		if ok {
			self.End = v
			self.present["end"] = true
			return nil
		} else {
			return fmt.Errorf("Field end/End: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "excludeMetadata", "ExcludeMetadata":
		v, ok := value.(bool)
		if ok {
			self.ExcludeMetadata = v
			self.present["excludeMetadata"] = true
			return nil
		} else {
			return fmt.Errorf("Field excludeMetadata/ExcludeMetadata: value %v(%T) couldn't be cast to type bool", value, value)
		}

	}
}

func (self *SingularityS3SearchRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityS3SearchRequest", name)

	case "listOnly", "ListOnly":
		if self.present != nil {
			if _, ok := self.present["listOnly"]; ok {
				return self.ListOnly, nil
			}
		}
		return nil, fmt.Errorf("Field ListOnly no set on ListOnly %+v", self)

	case "maxPerPage", "MaxPerPage":
		if self.present != nil {
			if _, ok := self.present["maxPerPage"]; ok {
				return self.MaxPerPage, nil
			}
		}
		return nil, fmt.Errorf("Field MaxPerPage no set on MaxPerPage %+v", self)

	case "requestsAndDeploys", "RequestsAndDeploys":
		if self.present != nil {
			if _, ok := self.present["requestsAndDeploys"]; ok {
				return self.RequestsAndDeploys, nil
			}
		}
		return nil, fmt.Errorf("Field RequestsAndDeploys no set on RequestsAndDeploys %+v", self)

	case "taskIds", "TaskIds":
		if self.present != nil {
			if _, ok := self.present["taskIds"]; ok {
				return self.TaskIds, nil
			}
		}
		return nil, fmt.Errorf("Field TaskIds no set on TaskIds %+v", self)

	case "start", "Start":
		if self.present != nil {
			if _, ok := self.present["start"]; ok {
				return self.Start, nil
			}
		}
		return nil, fmt.Errorf("Field Start no set on Start %+v", self)

	case "end", "End":
		if self.present != nil {
			if _, ok := self.present["end"]; ok {
				return self.End, nil
			}
		}
		return nil, fmt.Errorf("Field End no set on End %+v", self)

	case "excludeMetadata", "ExcludeMetadata":
		if self.present != nil {
			if _, ok := self.present["excludeMetadata"]; ok {
				return self.ExcludeMetadata, nil
			}
		}
		return nil, fmt.Errorf("Field ExcludeMetadata no set on ExcludeMetadata %+v", self)

	}
}

func (self *SingularityS3SearchRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityS3SearchRequest", name)

	case "listOnly", "ListOnly":
		self.present["listOnly"] = false

	case "maxPerPage", "MaxPerPage":
		self.present["maxPerPage"] = false

	case "requestsAndDeploys", "RequestsAndDeploys":
		self.present["requestsAndDeploys"] = false

	case "taskIds", "TaskIds":
		self.present["taskIds"] = false

	case "start", "Start":
		self.present["start"] = false

	case "end", "End":
		self.present["end"] = false

	case "excludeMetadata", "ExcludeMetadata":
		self.present["excludeMetadata"] = false

	}

	return nil
}

func (self *SingularityS3SearchRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityS3SearchRequestList []*SingularityS3SearchRequest

func (self *SingularityS3SearchRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityS3SearchRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityS3SearchRequestList cannot copy the values from %#v", other)
}

func (list *SingularityS3SearchRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityS3SearchRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityS3SearchRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
