package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskReconciliationStatistics struct {
	TaskReconciliationResponseCount  *int64   `json:"taskReconciliationResponseCount,omitempty"`
	TaskReconciliationResponseP999   *float64 `json:"taskReconciliationResponseP999,omitempty"`
	TaskReconciliationStartedAt      *int64   `json:"taskReconciliationStartedAt,omitempty"`
	TaskReconciliationResponseMax    *int64   `json:"taskReconciliationResponseMax,omitempty"`
	TaskReconciliationResponseP75    *float64 `json:"taskReconciliationResponseP75,omitempty"`
	TaskReconciliationResponseP99    *float64 `json:"taskReconciliationResponseP99,omitempty"`
	TaskReconciliationResponseStddev *float64 `json:"taskReconciliationResponseStddev,omitempty"`
	TaskReconciliationDurationMillis *int64   `json:"taskReconciliationDurationMillis,omitempty"`
	TaskReconciliationResponseMean   *float64 `json:"taskReconciliationResponseMean,omitempty"`
	TaskReconciliationResponseMin    *int64   `json:"taskReconciliationResponseMin,omitempty"`
	TaskReconciliationResponseP98    *float64 `json:"taskReconciliationResponseP98,omitempty"`
	TaskReconciliationIterations     *int32   `json:"taskReconciliationIterations,omitempty"`
	TaskReconciliationResponseP50    *float64 `json:"taskReconciliationResponseP50,omitempty"`
	TaskReconciliationResponseP95    *float64 `json:"taskReconciliationResponseP95,omitempty"`
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

func (self *SingularityTaskReconciliationStatistics) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskReconciliationStatistics) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskReconciliationStatistics) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskReconciliationStatistics", name)

	case "taskReconciliationResponseCount", "TaskReconciliationResponseCount":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationResponseCount = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseCount/TaskReconciliationResponseCount: value %v(%T) couldn't be cast to type int64", value, value)

	case "taskReconciliationResponseP999", "TaskReconciliationResponseP999":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP999 = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseP999/TaskReconciliationResponseP999: value %v(%T) couldn't be cast to type float64", value, value)

	case "taskReconciliationStartedAt", "TaskReconciliationStartedAt":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationStartedAt = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationStartedAt/TaskReconciliationStartedAt: value %v(%T) couldn't be cast to type int64", value, value)

	case "taskReconciliationResponseMax", "TaskReconciliationResponseMax":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationResponseMax = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseMax/TaskReconciliationResponseMax: value %v(%T) couldn't be cast to type int64", value, value)

	case "taskReconciliationResponseP75", "TaskReconciliationResponseP75":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP75 = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseP75/TaskReconciliationResponseP75: value %v(%T) couldn't be cast to type float64", value, value)

	case "taskReconciliationResponseP99", "TaskReconciliationResponseP99":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP99 = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseP99/TaskReconciliationResponseP99: value %v(%T) couldn't be cast to type float64", value, value)

	case "taskReconciliationResponseStddev", "TaskReconciliationResponseStddev":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseStddev = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseStddev/TaskReconciliationResponseStddev: value %v(%T) couldn't be cast to type float64", value, value)

	case "taskReconciliationDurationMillis", "TaskReconciliationDurationMillis":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationDurationMillis = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationDurationMillis/TaskReconciliationDurationMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "taskReconciliationResponseMean", "TaskReconciliationResponseMean":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseMean = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseMean/TaskReconciliationResponseMean: value %v(%T) couldn't be cast to type float64", value, value)

	case "taskReconciliationResponseMin", "TaskReconciliationResponseMin":
		v, ok := value.(int64)
		if ok {
			self.TaskReconciliationResponseMin = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseMin/TaskReconciliationResponseMin: value %v(%T) couldn't be cast to type int64", value, value)

	case "taskReconciliationResponseP98", "TaskReconciliationResponseP98":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP98 = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseP98/TaskReconciliationResponseP98: value %v(%T) couldn't be cast to type float64", value, value)

	case "taskReconciliationIterations", "TaskReconciliationIterations":
		v, ok := value.(int32)
		if ok {
			self.TaskReconciliationIterations = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationIterations/TaskReconciliationIterations: value %v(%T) couldn't be cast to type int32", value, value)

	case "taskReconciliationResponseP50", "TaskReconciliationResponseP50":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP50 = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseP50/TaskReconciliationResponseP50: value %v(%T) couldn't be cast to type float64", value, value)

	case "taskReconciliationResponseP95", "TaskReconciliationResponseP95":
		v, ok := value.(float64)
		if ok {
			self.TaskReconciliationResponseP95 = &v
			return nil
		}
		return fmt.Errorf("Field taskReconciliationResponseP95/TaskReconciliationResponseP95: value %v(%T) couldn't be cast to type float64", value, value)

	}
}

func (self *SingularityTaskReconciliationStatistics) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskReconciliationStatistics", name)

	case "taskReconciliationResponseCount", "TaskReconciliationResponseCount":
		return *self.TaskReconciliationResponseCount, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseCount no set on TaskReconciliationResponseCount %+v", self)

	case "taskReconciliationResponseP999", "TaskReconciliationResponseP999":
		return *self.TaskReconciliationResponseP999, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseP999 no set on TaskReconciliationResponseP999 %+v", self)

	case "taskReconciliationStartedAt", "TaskReconciliationStartedAt":
		return *self.TaskReconciliationStartedAt, nil
		return nil, fmt.Errorf("Field TaskReconciliationStartedAt no set on TaskReconciliationStartedAt %+v", self)

	case "taskReconciliationResponseMax", "TaskReconciliationResponseMax":
		return *self.TaskReconciliationResponseMax, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseMax no set on TaskReconciliationResponseMax %+v", self)

	case "taskReconciliationResponseP75", "TaskReconciliationResponseP75":
		return *self.TaskReconciliationResponseP75, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseP75 no set on TaskReconciliationResponseP75 %+v", self)

	case "taskReconciliationResponseP99", "TaskReconciliationResponseP99":
		return *self.TaskReconciliationResponseP99, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseP99 no set on TaskReconciliationResponseP99 %+v", self)

	case "taskReconciliationResponseStddev", "TaskReconciliationResponseStddev":
		return *self.TaskReconciliationResponseStddev, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseStddev no set on TaskReconciliationResponseStddev %+v", self)

	case "taskReconciliationDurationMillis", "TaskReconciliationDurationMillis":
		return *self.TaskReconciliationDurationMillis, nil
		return nil, fmt.Errorf("Field TaskReconciliationDurationMillis no set on TaskReconciliationDurationMillis %+v", self)

	case "taskReconciliationResponseMean", "TaskReconciliationResponseMean":
		return *self.TaskReconciliationResponseMean, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseMean no set on TaskReconciliationResponseMean %+v", self)

	case "taskReconciliationResponseMin", "TaskReconciliationResponseMin":
		return *self.TaskReconciliationResponseMin, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseMin no set on TaskReconciliationResponseMin %+v", self)

	case "taskReconciliationResponseP98", "TaskReconciliationResponseP98":
		return *self.TaskReconciliationResponseP98, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseP98 no set on TaskReconciliationResponseP98 %+v", self)

	case "taskReconciliationIterations", "TaskReconciliationIterations":
		return *self.TaskReconciliationIterations, nil
		return nil, fmt.Errorf("Field TaskReconciliationIterations no set on TaskReconciliationIterations %+v", self)

	case "taskReconciliationResponseP50", "TaskReconciliationResponseP50":
		return *self.TaskReconciliationResponseP50, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseP50 no set on TaskReconciliationResponseP50 %+v", self)

	case "taskReconciliationResponseP95", "TaskReconciliationResponseP95":
		return *self.TaskReconciliationResponseP95, nil
		return nil, fmt.Errorf("Field TaskReconciliationResponseP95 no set on TaskReconciliationResponseP95 %+v", self)

	}
}

func (self *SingularityTaskReconciliationStatistics) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskReconciliationStatistics", name)

	case "taskReconciliationResponseCount", "TaskReconciliationResponseCount":
		self.TaskReconciliationResponseCount = nil

	case "taskReconciliationResponseP999", "TaskReconciliationResponseP999":
		self.TaskReconciliationResponseP999 = nil

	case "taskReconciliationStartedAt", "TaskReconciliationStartedAt":
		self.TaskReconciliationStartedAt = nil

	case "taskReconciliationResponseMax", "TaskReconciliationResponseMax":
		self.TaskReconciliationResponseMax = nil

	case "taskReconciliationResponseP75", "TaskReconciliationResponseP75":
		self.TaskReconciliationResponseP75 = nil

	case "taskReconciliationResponseP99", "TaskReconciliationResponseP99":
		self.TaskReconciliationResponseP99 = nil

	case "taskReconciliationResponseStddev", "TaskReconciliationResponseStddev":
		self.TaskReconciliationResponseStddev = nil

	case "taskReconciliationDurationMillis", "TaskReconciliationDurationMillis":
		self.TaskReconciliationDurationMillis = nil

	case "taskReconciliationResponseMean", "TaskReconciliationResponseMean":
		self.TaskReconciliationResponseMean = nil

	case "taskReconciliationResponseMin", "TaskReconciliationResponseMin":
		self.TaskReconciliationResponseMin = nil

	case "taskReconciliationResponseP98", "TaskReconciliationResponseP98":
		self.TaskReconciliationResponseP98 = nil

	case "taskReconciliationIterations", "TaskReconciliationIterations":
		self.TaskReconciliationIterations = nil

	case "taskReconciliationResponseP50", "TaskReconciliationResponseP50":
		self.TaskReconciliationResponseP50 = nil

	case "taskReconciliationResponseP95", "TaskReconciliationResponseP95":
		self.TaskReconciliationResponseP95 = nil

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
