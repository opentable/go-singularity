package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskHistory struct {
	present map[string]bool

	Directory string `json:"directory,omitempty"`

	HealthcheckResults SingularityTaskHealthcheckResultList `json:"healthcheckResults"`

	LoadBalancerUpdates SingularityLoadBalancerUpdateList `json:"loadBalancerUpdates"`

	ShellCommandHistory SingularityTaskShellCommandHistoryList `json:"shellCommandHistory"`

	Task *SingularityTask `json:"task"`

	TaskUpdates SingularityTaskHistoryUpdateList `json:"taskUpdates"`
}

func (self *SingularityTaskHistory) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskHistory) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskHistory); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskHistory cannot copy the values from %#v", other)
}

func (self *SingularityTaskHistory) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskHistory) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskHistory) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskHistory) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskHistory) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHistory", name)

	case "directory", "Directory":
		v, ok := value.(string)
		if ok {
			self.Directory = v
			self.present["directory"] = true
			return nil
		} else {
			return fmt.Errorf("Field directory/Directory: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "healthcheckResults", "HealthcheckResults":
		v, ok := value.(SingularityTaskHealthcheckResultList)
		if ok {
			self.HealthcheckResults = v
			self.present["healthcheckResults"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthcheckResults/HealthcheckResults: value %v(%T) couldn't be cast to type SingularityTaskHealthcheckResultList", value, value)
		}

	case "loadBalancerUpdates", "LoadBalancerUpdates":
		v, ok := value.(SingularityLoadBalancerUpdateList)
		if ok {
			self.LoadBalancerUpdates = v
			self.present["loadBalancerUpdates"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalancerUpdates/LoadBalancerUpdates: value %v(%T) couldn't be cast to type SingularityLoadBalancerUpdateList", value, value)
		}

	case "shellCommandHistory", "ShellCommandHistory":
		v, ok := value.(SingularityTaskShellCommandHistoryList)
		if ok {
			self.ShellCommandHistory = v
			self.present["shellCommandHistory"] = true
			return nil
		} else {
			return fmt.Errorf("Field shellCommandHistory/ShellCommandHistory: value %v(%T) couldn't be cast to type SingularityTaskShellCommandHistoryList", value, value)
		}

	case "task", "Task":
		v, ok := value.(*SingularityTask)
		if ok {
			self.Task = v
			self.present["task"] = true
			return nil
		} else {
			return fmt.Errorf("Field task/Task: value %v(%T) couldn't be cast to type *SingularityTask", value, value)
		}

	case "taskUpdates", "TaskUpdates":
		v, ok := value.(SingularityTaskHistoryUpdateList)
		if ok {
			self.TaskUpdates = v
			self.present["taskUpdates"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskUpdates/TaskUpdates: value %v(%T) couldn't be cast to type SingularityTaskHistoryUpdateList", value, value)
		}

	}
}

func (self *SingularityTaskHistory) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskHistory", name)

	case "directory", "Directory":
		if self.present != nil {
			if _, ok := self.present["directory"]; ok {
				return self.Directory, nil
			}
		}
		return nil, fmt.Errorf("Field Directory no set on Directory %+v", self)

	case "healthcheckResults", "HealthcheckResults":
		if self.present != nil {
			if _, ok := self.present["healthcheckResults"]; ok {
				return self.HealthcheckResults, nil
			}
		}
		return nil, fmt.Errorf("Field HealthcheckResults no set on HealthcheckResults %+v", self)

	case "loadBalancerUpdates", "LoadBalancerUpdates":
		if self.present != nil {
			if _, ok := self.present["loadBalancerUpdates"]; ok {
				return self.LoadBalancerUpdates, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalancerUpdates no set on LoadBalancerUpdates %+v", self)

	case "shellCommandHistory", "ShellCommandHistory":
		if self.present != nil {
			if _, ok := self.present["shellCommandHistory"]; ok {
				return self.ShellCommandHistory, nil
			}
		}
		return nil, fmt.Errorf("Field ShellCommandHistory no set on ShellCommandHistory %+v", self)

	case "task", "Task":
		if self.present != nil {
			if _, ok := self.present["task"]; ok {
				return self.Task, nil
			}
		}
		return nil, fmt.Errorf("Field Task no set on Task %+v", self)

	case "taskUpdates", "TaskUpdates":
		if self.present != nil {
			if _, ok := self.present["taskUpdates"]; ok {
				return self.TaskUpdates, nil
			}
		}
		return nil, fmt.Errorf("Field TaskUpdates no set on TaskUpdates %+v", self)

	}
}

func (self *SingularityTaskHistory) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHistory", name)

	case "directory", "Directory":
		self.present["directory"] = false

	case "healthcheckResults", "HealthcheckResults":
		self.present["healthcheckResults"] = false

	case "loadBalancerUpdates", "LoadBalancerUpdates":
		self.present["loadBalancerUpdates"] = false

	case "shellCommandHistory", "ShellCommandHistory":
		self.present["shellCommandHistory"] = false

	case "task", "Task":
		self.present["task"] = false

	case "taskUpdates", "TaskUpdates":
		self.present["taskUpdates"] = false

	}

	return nil
}

func (self *SingularityTaskHistory) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskHistoryList []*SingularityTaskHistory

func (self *SingularityTaskHistoryList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskHistoryList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskHistoryList cannot copy the values from %#v", other)
}

func (list *SingularityTaskHistoryList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskHistoryList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskHistoryList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
