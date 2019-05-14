package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityS3SearchRequest struct {
	RequestsAndDeploys *map[string]swaggering.StringList `json:"requestsAndDeploys,omitempty"`
	TaskIds            *swaggering.StringList            `json:"taskIds,omitempty"`
	Start              *int64                            `json:"start,omitempty"`
	End                *int64                            `json:"end,omitempty"`
	ExcludeMetadata    *bool                             `json:"excludeMetadata,omitempty"`
	ListOnly           *bool                             `json:"listOnly,omitempty"`
	MaxPerPage         *int32                            `json:"maxPerPage,omitempty"`
	// Invalid field: ContinuationTokens *notfound.Map[string,ContinuationToken] `json:"continuationTokens,omitempty"`

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

func (self *SingularityS3SearchRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityS3SearchRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityS3SearchRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityS3SearchRequest", name)

	case "requestsAndDeploys", "RequestsAndDeploys":
		v, ok := value.(map[string]swaggering.StringList)
		if ok {
			self.RequestsAndDeploys = &v
			return nil
		}
		return fmt.Errorf("Field requestsAndDeploys/RequestsAndDeploys: value %v(%T) couldn't be cast to type map[string]swaggering.StringList", value, value)

	case "taskIds", "TaskIds":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.TaskIds = &v
			return nil
		}
		return fmt.Errorf("Field taskIds/TaskIds: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "start", "Start":
		v, ok := value.(int64)
		if ok {
			self.Start = &v
			return nil
		}
		return fmt.Errorf("Field start/Start: value %v(%T) couldn't be cast to type int64", value, value)

	case "end", "End":
		v, ok := value.(int64)
		if ok {
			self.End = &v
			return nil
		}
		return fmt.Errorf("Field end/End: value %v(%T) couldn't be cast to type int64", value, value)

	case "excludeMetadata", "ExcludeMetadata":
		v, ok := value.(bool)
		if ok {
			self.ExcludeMetadata = &v
			return nil
		}
		return fmt.Errorf("Field excludeMetadata/ExcludeMetadata: value %v(%T) couldn't be cast to type bool", value, value)

	case "listOnly", "ListOnly":
		v, ok := value.(bool)
		if ok {
			self.ListOnly = &v
			return nil
		}
		return fmt.Errorf("Field listOnly/ListOnly: value %v(%T) couldn't be cast to type bool", value, value)

	case "maxPerPage", "MaxPerPage":
		v, ok := value.(int32)
		if ok {
			self.MaxPerPage = &v
			return nil
		}
		return fmt.Errorf("Field maxPerPage/MaxPerPage: value %v(%T) couldn't be cast to type int32", value, value)

	}
}

func (self *SingularityS3SearchRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityS3SearchRequest", name)

	case "requestsAndDeploys", "RequestsAndDeploys":
		return *self.RequestsAndDeploys, nil
		return nil, fmt.Errorf("Field RequestsAndDeploys no set on RequestsAndDeploys %+v", self)

	case "taskIds", "TaskIds":
		return *self.TaskIds, nil
		return nil, fmt.Errorf("Field TaskIds no set on TaskIds %+v", self)

	case "start", "Start":
		return *self.Start, nil
		return nil, fmt.Errorf("Field Start no set on Start %+v", self)

	case "end", "End":
		return *self.End, nil
		return nil, fmt.Errorf("Field End no set on End %+v", self)

	case "excludeMetadata", "ExcludeMetadata":
		return *self.ExcludeMetadata, nil
		return nil, fmt.Errorf("Field ExcludeMetadata no set on ExcludeMetadata %+v", self)

	case "listOnly", "ListOnly":
		return *self.ListOnly, nil
		return nil, fmt.Errorf("Field ListOnly no set on ListOnly %+v", self)

	case "maxPerPage", "MaxPerPage":
		return *self.MaxPerPage, nil
		return nil, fmt.Errorf("Field MaxPerPage no set on MaxPerPage %+v", self)

	}
}

func (self *SingularityS3SearchRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityS3SearchRequest", name)

	case "requestsAndDeploys", "RequestsAndDeploys":
		self.RequestsAndDeploys = nil

	case "taskIds", "TaskIds":
		self.TaskIds = nil

	case "start", "Start":
		self.Start = nil

	case "end", "End":
		self.End = nil

	case "excludeMetadata", "ExcludeMetadata":
		self.ExcludeMetadata = nil

	case "listOnly", "ListOnly":
		self.ListOnly = nil

	case "maxPerPage", "MaxPerPage":
		self.MaxPerPage = nil

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
