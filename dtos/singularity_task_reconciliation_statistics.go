package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskReconciliationStatistics struct {
	present map[string]bool

	TaskReconciliationStartedAt int64 `json:"taskReconciliationStartedAt"`

	TaskReconciliationDurationMillis int64 `json:"taskReconciliationDurationMillis"`

	TaskReconciliationIterations int32 `json:"taskReconciliationIterations"`

	TaskReconciliationResponseCount int64 `json:"taskReconciliationResponseCount"`

	TaskReconciliationResponseP95 float64 `json:"taskReconciliationResponseP95"`

	TaskReconciliationResponseP99 float64 `json:"taskReconciliationResponseP99"`

	TaskReconciliationResponseP50 float64 `json:"taskReconciliationResponseP50"`

	TaskReconciliationResponseP75 float64 `json:"taskReconciliationResponseP75"`

	TaskReconciliationResponseP999 float64 `json:"taskReconciliationResponseP999"`

	TaskReconciliationResponseP98 float64 `json:"taskReconciliationResponseP98"`

	TaskReconciliationResponseStddev float64 `json:"taskReconciliationResponseStddev"`

	TaskReconciliationResponseMax int64 `json:"taskReconciliationResponseMax"`

	TaskReconciliationResponseMean float64 `json:"taskReconciliationResponseMean"`

	TaskReconciliationResponseMin int64 `json:"taskReconciliationResponseMin"`
}

func (self *SingularityTaskReconciliationStatistics) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskReconciliationStatistics) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskReconciliationStatistics); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskReconciliationStatistics cannot copy the values from %#v", other)
}

func (self *SingularityTaskReconciliationStatistics) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskReconciliationStatistics) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskReconciliationStatistics) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskReconciliationStatistics) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskReconciliationStatistics) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskReconciliationStatistics", name)

	case "taskReconciliationStartedAt", "TaskReconciliationStartedAt":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationStartedAt = v
			self.present["taskReconciliationStartedAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationStartedAt/TaskReconciliationStartedAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "taskReconciliationDurationMillis", "TaskReconciliationDurationMillis":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationDurationMillis = v
			self.present["taskReconciliationDurationMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationDurationMillis/TaskReconciliationDurationMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "taskReconciliationIterations", "TaskReconciliationIterations":
		v, ok := value.(int32)
		if ok {
			self.TaskReconciliationIterations = v
			self.present["taskReconciliationIterations"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationIterations/TaskReconciliationIterations: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "taskReconciliationResponseCount", "TaskReconciliationResponseCount":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationResponseCount = v
			self.present["taskReconciliationResponseCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseCount/TaskReconciliationResponseCount: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "taskReconciliationResponseP95", "TaskReconciliationResponseP95":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP95 = v
			self.present["taskReconciliationResponseP95"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseP95/TaskReconciliationResponseP95: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "taskReconciliationResponseP99", "TaskReconciliationResponseP99":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP99 = v
			self.present["taskReconciliationResponseP99"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseP99/TaskReconciliationResponseP99: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "taskReconciliationResponseP50", "TaskReconciliationResponseP50":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP50 = v
			self.present["taskReconciliationResponseP50"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseP50/TaskReconciliationResponseP50: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "taskReconciliationResponseP75", "TaskReconciliationResponseP75":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP75 = v
			self.present["taskReconciliationResponseP75"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseP75/TaskReconciliationResponseP75: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "taskReconciliationResponseP999", "TaskReconciliationResponseP999":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP999 = v
			self.present["taskReconciliationResponseP999"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseP999/TaskReconciliationResponseP999: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "taskReconciliationResponseP98", "TaskReconciliationResponseP98":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP98 = v
			self.present["taskReconciliationResponseP98"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseP98/TaskReconciliationResponseP98: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "taskReconciliationResponseStddev", "TaskReconciliationResponseStddev":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseStddev = v
			self.present["taskReconciliationResponseStddev"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseStddev/TaskReconciliationResponseStddev: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "taskReconciliationResponseMax", "TaskReconciliationResponseMax":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationResponseMax = v
			self.present["taskReconciliationResponseMax"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseMax/TaskReconciliationResponseMax: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "taskReconciliationResponseMean", "TaskReconciliationResponseMean":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseMean = v
			self.present["taskReconciliationResponseMean"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseMean/TaskReconciliationResponseMean: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "taskReconciliationResponseMin", "TaskReconciliationResponseMin":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationResponseMin = v
			self.present["taskReconciliationResponseMin"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskReconciliationResponseMin/TaskReconciliationResponseMin: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityTaskReconciliationStatistics) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskReconciliationStatistics", name)

	case "taskReconciliationStartedAt", "TaskReconciliationStartedAt":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationStartedAt"]; ok {
				return self.TaskReconciliationStartedAt, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationStartedAt no set on TaskReconciliationStartedAt %+v", self)

	case "taskReconciliationDurationMillis", "TaskReconciliationDurationMillis":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationDurationMillis"]; ok {
				return self.TaskReconciliationDurationMillis, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationDurationMillis no set on TaskReconciliationDurationMillis %+v", self)

	case "taskReconciliationIterations", "TaskReconciliationIterations":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationIterations"]; ok {
				return self.TaskReconciliationIterations, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationIterations no set on TaskReconciliationIterations %+v", self)

	case "taskReconciliationResponseCount", "TaskReconciliationResponseCount":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseCount"]; ok {
				return self.TaskReconciliationResponseCount, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseCount no set on TaskReconciliationResponseCount %+v", self)

	case "taskReconciliationResponseP95", "TaskReconciliationResponseP95":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseP95"]; ok {
				return self.TaskReconciliationResponseP95, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseP95 no set on TaskReconciliationResponseP95 %+v", self)

	case "taskReconciliationResponseP99", "TaskReconciliationResponseP99":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseP99"]; ok {
				return self.TaskReconciliationResponseP99, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseP99 no set on TaskReconciliationResponseP99 %+v", self)

	case "taskReconciliationResponseP50", "TaskReconciliationResponseP50":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseP50"]; ok {
				return self.TaskReconciliationResponseP50, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseP50 no set on TaskReconciliationResponseP50 %+v", self)

	case "taskReconciliationResponseP75", "TaskReconciliationResponseP75":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseP75"]; ok {
				return self.TaskReconciliationResponseP75, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseP75 no set on TaskReconciliationResponseP75 %+v", self)

	case "taskReconciliationResponseP999", "TaskReconciliationResponseP999":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseP999"]; ok {
				return self.TaskReconciliationResponseP999, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseP999 no set on TaskReconciliationResponseP999 %+v", self)

	case "taskReconciliationResponseP98", "TaskReconciliationResponseP98":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseP98"]; ok {
				return self.TaskReconciliationResponseP98, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseP98 no set on TaskReconciliationResponseP98 %+v", self)

	case "taskReconciliationResponseStddev", "TaskReconciliationResponseStddev":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseStddev"]; ok {
				return self.TaskReconciliationResponseStddev, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseStddev no set on TaskReconciliationResponseStddev %+v", self)

	case "taskReconciliationResponseMax", "TaskReconciliationResponseMax":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseMax"]; ok {
				return self.TaskReconciliationResponseMax, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseMax no set on TaskReconciliationResponseMax %+v", self)

	case "taskReconciliationResponseMean", "TaskReconciliationResponseMean":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseMean"]; ok {
				return self.TaskReconciliationResponseMean, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseMean no set on TaskReconciliationResponseMean %+v", self)

	case "taskReconciliationResponseMin", "TaskReconciliationResponseMin":
		if self.present != nil {
			if _, ok := self.present["taskReconciliationResponseMin"]; ok {
				return self.TaskReconciliationResponseMin, nil
			}
		}
		return nil, fmt.Errorf("Field TaskReconciliationResponseMin no set on TaskReconciliationResponseMin %+v", self)

	}
}

func (self *SingularityTaskReconciliationStatistics) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskReconciliationStatistics", name)

	case "taskReconciliationStartedAt", "TaskReconciliationStartedAt":
		self.present["taskReconciliationStartedAt"] = false

	case "taskReconciliationDurationMillis", "TaskReconciliationDurationMillis":
		self.present["taskReconciliationDurationMillis"] = false

	case "taskReconciliationIterations", "TaskReconciliationIterations":
		self.present["taskReconciliationIterations"] = false

	case "taskReconciliationResponseCount", "TaskReconciliationResponseCount":
		self.present["taskReconciliationResponseCount"] = false

	case "taskReconciliationResponseP95", "TaskReconciliationResponseP95":
		self.present["taskReconciliationResponseP95"] = false

	case "taskReconciliationResponseP99", "TaskReconciliationResponseP99":
		self.present["taskReconciliationResponseP99"] = false

	case "taskReconciliationResponseP50", "TaskReconciliationResponseP50":
		self.present["taskReconciliationResponseP50"] = false

	case "taskReconciliationResponseP75", "TaskReconciliationResponseP75":
		self.present["taskReconciliationResponseP75"] = false

	case "taskReconciliationResponseP999", "TaskReconciliationResponseP999":
		self.present["taskReconciliationResponseP999"] = false

	case "taskReconciliationResponseP98", "TaskReconciliationResponseP98":
		self.present["taskReconciliationResponseP98"] = false

	case "taskReconciliationResponseStddev", "TaskReconciliationResponseStddev":
		self.present["taskReconciliationResponseStddev"] = false

	case "taskReconciliationResponseMax", "TaskReconciliationResponseMax":
		self.present["taskReconciliationResponseMax"] = false

	case "taskReconciliationResponseMean", "TaskReconciliationResponseMean":
		self.present["taskReconciliationResponseMean"] = false

	case "taskReconciliationResponseMin", "TaskReconciliationResponseMin":
		self.present["taskReconciliationResponseMin"] = false

	}

	return nil
}

func (self *SingularityTaskReconciliationStatistics) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskReconciliationStatisticsList []*SingularityTaskReconciliationStatistics

func (self *SingularityTaskReconciliationStatisticsList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskReconciliationStatisticsList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskReconciliationStatisticsList cannot copy the values from %#v", other)
}

func (list *SingularityTaskReconciliationStatisticsList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskReconciliationStatisticsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskReconciliationStatisticsList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
