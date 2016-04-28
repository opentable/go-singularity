package dtos

import (
	"fmt"
	"io"
)

type SingularityTaskHealthcheckResult struct {
	present        map[string]bool
	DurationMillis int64              `json:"durationMillis"`
	ErrorMessage   string             `json:"errorMessage,omitempty"`
	ResponseBody   string             `json:"responseBody,omitempty"`
	StatusCode     int32              `json:"statusCode"`
	TaskId         *SingularityTaskId `json:"taskId"`
	Timestamp      int64              `json:"timestamp"`
}

func (self *SingularityTaskHealthcheckResult) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskHealthcheckResult) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *SingularityTaskHealthcheckResult) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskHealthcheckResult) FormatJSON() string {
	return FormatJSON(self)
}

func (self *SingularityTaskHealthcheckResult) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *SingularityTaskHealthcheckResult) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHealthcheckResult", name)

	case "durationMillis", "DurationMillis":
		v, ok := value.(int64)
		if ok {
			self.DurationMillis = v
			self.present["durationMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field durationMillis/DurationMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "errorMessage", "ErrorMessage":
		v, ok := value.(string)
		if ok {
			self.ErrorMessage = v
			self.present["errorMessage"] = true
			return nil
		} else {
			return fmt.Errorf("Field errorMessage/ErrorMessage: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "responseBody", "ResponseBody":
		v, ok := value.(string)
		if ok {
			self.ResponseBody = v
			self.present["responseBody"] = true
			return nil
		} else {
			return fmt.Errorf("Field responseBody/ResponseBody: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "statusCode", "StatusCode":
		v, ok := value.(int32)
		if ok {
			self.StatusCode = v
			self.present["statusCode"] = true
			return nil
		} else {
			return fmt.Errorf("Field statusCode/StatusCode: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			self.present["taskId"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)
		}

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityTaskHealthcheckResult) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskHealthcheckResult", name)

	case "durationMillis", "DurationMillis":
		if self.present != nil {
			if _, ok := self.present["durationMillis"]; ok {
				return self.DurationMillis, nil
			}
		}
		return nil, fmt.Errorf("Field DurationMillis no set on DurationMillis %+v", self)

	case "errorMessage", "ErrorMessage":
		if self.present != nil {
			if _, ok := self.present["errorMessage"]; ok {
				return self.ErrorMessage, nil
			}
		}
		return nil, fmt.Errorf("Field ErrorMessage no set on ErrorMessage %+v", self)

	case "responseBody", "ResponseBody":
		if self.present != nil {
			if _, ok := self.present["responseBody"]; ok {
				return self.ResponseBody, nil
			}
		}
		return nil, fmt.Errorf("Field ResponseBody no set on ResponseBody %+v", self)

	case "statusCode", "StatusCode":
		if self.present != nil {
			if _, ok := self.present["statusCode"]; ok {
				return self.StatusCode, nil
			}
		}
		return nil, fmt.Errorf("Field StatusCode no set on StatusCode %+v", self)

	case "taskId", "TaskId":
		if self.present != nil {
			if _, ok := self.present["taskId"]; ok {
				return self.TaskId, nil
			}
		}
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityTaskHealthcheckResult) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHealthcheckResult", name)

	case "durationMillis", "DurationMillis":
		self.present["durationMillis"] = false

	case "errorMessage", "ErrorMessage":
		self.present["errorMessage"] = false

	case "responseBody", "ResponseBody":
		self.present["responseBody"] = false

	case "statusCode", "StatusCode":
		self.present["statusCode"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	}

	return nil
}

func (self *SingularityTaskHealthcheckResult) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type SingularityTaskHealthcheckResultList []*SingularityTaskHealthcheckResult

func (list *SingularityTaskHealthcheckResultList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
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
	return FormatJSON(list)
}
