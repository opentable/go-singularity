package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeployStatisticsExtendedTaskState string

const (
	SingularityDeployStatisticsExtendedTaskStateTASK_LAUNCHED        SingularityDeployStatisticsExtendedTaskState = "TASK_LAUNCHED"
	SingularityDeployStatisticsExtendedTaskStateTASK_STAGING         SingularityDeployStatisticsExtendedTaskState = "TASK_STAGING"
	SingularityDeployStatisticsExtendedTaskStateTASK_STARTING        SingularityDeployStatisticsExtendedTaskState = "TASK_STARTING"
	SingularityDeployStatisticsExtendedTaskStateTASK_RUNNING         SingularityDeployStatisticsExtendedTaskState = "TASK_RUNNING"
	SingularityDeployStatisticsExtendedTaskStateTASK_CLEANING        SingularityDeployStatisticsExtendedTaskState = "TASK_CLEANING"
	SingularityDeployStatisticsExtendedTaskStateTASK_KILLING         SingularityDeployStatisticsExtendedTaskState = "TASK_KILLING"
	SingularityDeployStatisticsExtendedTaskStateTASK_FINISHED        SingularityDeployStatisticsExtendedTaskState = "TASK_FINISHED"
	SingularityDeployStatisticsExtendedTaskStateTASK_FAILED          SingularityDeployStatisticsExtendedTaskState = "TASK_FAILED"
	SingularityDeployStatisticsExtendedTaskStateTASK_KILLED          SingularityDeployStatisticsExtendedTaskState = "TASK_KILLED"
	SingularityDeployStatisticsExtendedTaskStateTASK_LOST            SingularityDeployStatisticsExtendedTaskState = "TASK_LOST"
	SingularityDeployStatisticsExtendedTaskStateTASK_LOST_WHILE_DOWN SingularityDeployStatisticsExtendedTaskState = "TASK_LOST_WHILE_DOWN"
	SingularityDeployStatisticsExtendedTaskStateTASK_ERROR           SingularityDeployStatisticsExtendedTaskState = "TASK_ERROR"
)

type SingularityDeployStatistics struct {
	present map[string]bool

	RequestId string `json:"requestId,omitempty"`

	NumTasks int32 `json:"numTasks"`

	NumSequentialRetries int32 `json:"numSequentialRetries"`

	InstanceSequentialFailureTimestamps map[int64][]int64 `json:"instanceSequentialFailureTimestamps"`

	LastTaskState SingularityDeployStatisticsExtendedTaskState `json:"lastTaskState"`

	AverageRuntimeMillis int64 `json:"averageRuntimeMillis"`

	DeployId string `json:"deployId,omitempty"`

	NumSuccess int32 `json:"numSuccess"`

	NumFailures int32 `json:"numFailures"`

	LastFinishAt int64 `json:"lastFinishAt"`
}

func (self *SingularityDeployStatistics) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployStatistics) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployStatistics); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployStatistics cannot copy the values from %#v", other)
}

func (self *SingularityDeployStatistics) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDeployStatistics) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployStatistics) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployStatistics) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDeployStatistics) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployStatistics", name)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = v
			self.present["requestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "numTasks", "NumTasks":
		v, ok := value.(int32)
		if ok {
			self.NumTasks = v
			self.present["numTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field numTasks/NumTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "numSequentialRetries", "NumSequentialRetries":
		v, ok := value.(int32)
		if ok {
			self.NumSequentialRetries = v
			self.present["numSequentialRetries"] = true
			return nil
		} else {
			return fmt.Errorf("Field numSequentialRetries/NumSequentialRetries: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "instanceSequentialFailureTimestamps", "InstanceSequentialFailureTimestamps":
		v, ok := value.(map[int64][]int64)
		if ok {
			self.InstanceSequentialFailureTimestamps = v
			self.present["instanceSequentialFailureTimestamps"] = true
			return nil
		} else {
			return fmt.Errorf("Field instanceSequentialFailureTimestamps/InstanceSequentialFailureTimestamps: value %v(%T) couldn't be cast to type map[int64][]int64", value, value)
		}

	case "lastTaskState", "LastTaskState":
		v, ok := value.(SingularityDeployStatisticsExtendedTaskState)
		if ok {
			self.LastTaskState = v
			self.present["lastTaskState"] = true
			return nil
		} else {
			return fmt.Errorf("Field lastTaskState/LastTaskState: value %v(%T) couldn't be cast to type SingularityDeployStatisticsExtendedTaskState", value, value)
		}

	case "averageRuntimeMillis", "AverageRuntimeMillis":
		v, ok := value.(int64)
		if ok {
			self.AverageRuntimeMillis = v
			self.present["averageRuntimeMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field averageRuntimeMillis/AverageRuntimeMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = v
			self.present["deployId"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "numSuccess", "NumSuccess":
		v, ok := value.(int32)
		if ok {
			self.NumSuccess = v
			self.present["numSuccess"] = true
			return nil
		} else {
			return fmt.Errorf("Field numSuccess/NumSuccess: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "numFailures", "NumFailures":
		v, ok := value.(int32)
		if ok {
			self.NumFailures = v
			self.present["numFailures"] = true
			return nil
		} else {
			return fmt.Errorf("Field numFailures/NumFailures: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "lastFinishAt", "LastFinishAt":
		v, ok := value.(int64)
		if ok {
			self.LastFinishAt = v
			self.present["lastFinishAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field lastFinishAt/LastFinishAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityDeployStatistics) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployStatistics", name)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "numTasks", "NumTasks":
		if self.present != nil {
			if _, ok := self.present["numTasks"]; ok {
				return self.NumTasks, nil
			}
		}
		return nil, fmt.Errorf("Field NumTasks no set on NumTasks %+v", self)

	case "numSequentialRetries", "NumSequentialRetries":
		if self.present != nil {
			if _, ok := self.present["numSequentialRetries"]; ok {
				return self.NumSequentialRetries, nil
			}
		}
		return nil, fmt.Errorf("Field NumSequentialRetries no set on NumSequentialRetries %+v", self)

	case "instanceSequentialFailureTimestamps", "InstanceSequentialFailureTimestamps":
		if self.present != nil {
			if _, ok := self.present["instanceSequentialFailureTimestamps"]; ok {
				return self.InstanceSequentialFailureTimestamps, nil
			}
		}
		return nil, fmt.Errorf("Field InstanceSequentialFailureTimestamps no set on InstanceSequentialFailureTimestamps %+v", self)

	case "lastTaskState", "LastTaskState":
		if self.present != nil {
			if _, ok := self.present["lastTaskState"]; ok {
				return self.LastTaskState, nil
			}
		}
		return nil, fmt.Errorf("Field LastTaskState no set on LastTaskState %+v", self)

	case "averageRuntimeMillis", "AverageRuntimeMillis":
		if self.present != nil {
			if _, ok := self.present["averageRuntimeMillis"]; ok {
				return self.AverageRuntimeMillis, nil
			}
		}
		return nil, fmt.Errorf("Field AverageRuntimeMillis no set on AverageRuntimeMillis %+v", self)

	case "deployId", "DeployId":
		if self.present != nil {
			if _, ok := self.present["deployId"]; ok {
				return self.DeployId, nil
			}
		}
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "numSuccess", "NumSuccess":
		if self.present != nil {
			if _, ok := self.present["numSuccess"]; ok {
				return self.NumSuccess, nil
			}
		}
		return nil, fmt.Errorf("Field NumSuccess no set on NumSuccess %+v", self)

	case "numFailures", "NumFailures":
		if self.present != nil {
			if _, ok := self.present["numFailures"]; ok {
				return self.NumFailures, nil
			}
		}
		return nil, fmt.Errorf("Field NumFailures no set on NumFailures %+v", self)

	case "lastFinishAt", "LastFinishAt":
		if self.present != nil {
			if _, ok := self.present["lastFinishAt"]; ok {
				return self.LastFinishAt, nil
			}
		}
		return nil, fmt.Errorf("Field LastFinishAt no set on LastFinishAt %+v", self)

	}
}

func (self *SingularityDeployStatistics) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployStatistics", name)

	case "requestId", "RequestId":
		self.present["requestId"] = false

	case "numTasks", "NumTasks":
		self.present["numTasks"] = false

	case "numSequentialRetries", "NumSequentialRetries":
		self.present["numSequentialRetries"] = false

	case "instanceSequentialFailureTimestamps", "InstanceSequentialFailureTimestamps":
		self.present["instanceSequentialFailureTimestamps"] = false

	case "lastTaskState", "LastTaskState":
		self.present["lastTaskState"] = false

	case "averageRuntimeMillis", "AverageRuntimeMillis":
		self.present["averageRuntimeMillis"] = false

	case "deployId", "DeployId":
		self.present["deployId"] = false

	case "numSuccess", "NumSuccess":
		self.present["numSuccess"] = false

	case "numFailures", "NumFailures":
		self.present["numFailures"] = false

	case "lastFinishAt", "LastFinishAt":
		self.present["lastFinishAt"] = false

	}

	return nil
}

func (self *SingularityDeployStatistics) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployStatisticsList []*SingularityDeployStatistics

func (self *SingularityDeployStatisticsList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployStatisticsList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployStatisticsList cannot copy the values from %#v", other)
}

func (list *SingularityDeployStatisticsList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployStatisticsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployStatisticsList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
