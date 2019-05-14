package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskHealthcheckResult struct {
	ErrorMessage   *string            `json:"errorMessage,omitempty"`
	Startup        *bool              `json:"startup,omitempty"`
	Timestamp      *int64             `json:"timestamp,omitempty"`
	TaskId         *SingularityTaskId `json:"taskId,omitempty"`
	StatusCode     *int32             `json:"statusCode,omitempty"`
	DurationMillis *int64             `json:"durationMillis,omitempty"`
	ResponseBody   *string            `json:"responseBody,omitempty"`
}

func (self *SingularityTaskHealthcheckResult) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskHealthcheckResult) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskHealthcheckResult); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskHealthcheckResult cannot copy the values from %#v", other)
}

func (self *SingularityTaskHealthcheckResult) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskHealthcheckResult) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskHealthcheckResult) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHealthcheckResult", name)

	case "errorMessage", "ErrorMessage":
		v, ok := value.(string)
		if ok {
			self.ErrorMessage = &v
			return nil
		}
		return fmt.Errorf("Field errorMessage/ErrorMessage: value %v(%T) couldn't be cast to type string", value, value)

	case "startup", "Startup":
		v, ok := value.(bool)
		if ok {
			self.Startup = &v
			return nil
		}
		return fmt.Errorf("Field startup/Startup: value %v(%T) couldn't be cast to type bool", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			return nil
		}
		return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)

	case "statusCode", "StatusCode":
		v, ok := value.(int32)
		if ok {
			self.StatusCode = &v
			return nil
		}
		return fmt.Errorf("Field statusCode/StatusCode: value %v(%T) couldn't be cast to type int32", value, value)

	case "durationMillis", "DurationMillis":
		v, ok := value.(int64)
		if ok {
			self.DurationMillis = &v
			return nil
		}
		return fmt.Errorf("Field durationMillis/DurationMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "responseBody", "ResponseBody":
		v, ok := value.(string)
		if ok {
			self.ResponseBody = &v
			return nil
		}
		return fmt.Errorf("Field responseBody/ResponseBody: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityTaskHealthcheckResult) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskHealthcheckResult", name)

	case "errorMessage", "ErrorMessage":
		return *self.ErrorMessage, nil
		return nil, fmt.Errorf("Field ErrorMessage no set on ErrorMessage %+v", self)

	case "startup", "Startup":
		return *self.Startup, nil
		return nil, fmt.Errorf("Field Startup no set on Startup %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "taskId", "TaskId":
		return self.TaskId, nil
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "statusCode", "StatusCode":
		return *self.StatusCode, nil
		return nil, fmt.Errorf("Field StatusCode no set on StatusCode %+v", self)

	case "durationMillis", "DurationMillis":
		return *self.DurationMillis, nil
		return nil, fmt.Errorf("Field DurationMillis no set on DurationMillis %+v", self)

	case "responseBody", "ResponseBody":
		return *self.ResponseBody, nil
		return nil, fmt.Errorf("Field ResponseBody no set on ResponseBody %+v", self)

	}
}

func (self *SingularityTaskHealthcheckResult) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHealthcheckResult", name)

	case "errorMessage", "ErrorMessage":
		self.ErrorMessage = nil

	case "startup", "Startup":
		self.Startup = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "taskId", "TaskId":
		self.TaskId = nil

	case "statusCode", "StatusCode":
		self.StatusCode = nil

	case "durationMillis", "DurationMillis":
		self.DurationMillis = nil

	case "responseBody", "ResponseBody":
		self.ResponseBody = nil

	}

	return nil
}

func (self *SingularityTaskHealthcheckResult) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskHealthcheckResultList []*SingularityTaskHealthcheckResult

func (self *SingularityTaskHealthcheckResultList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskHealthcheckResultList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskHealthcheckResultList cannot copy the values from %#v", other)
}

func (list *SingularityTaskHealthcheckResultList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskHealthcheckResultList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskHealthcheckResultList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
