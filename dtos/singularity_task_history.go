package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskHistory struct {
	ShellCommandHistory *SingularityTaskShellCommandHistoryList `json:"shellCommandHistory,omitempty"`
	TaskMetadata        *SingularityTaskMetadataList            `json:"taskMetadata,omitempty"`
	TaskUpdates         *SingularityTaskHistoryUpdateList       `json:"taskUpdates,omitempty"`
	Directory           *string                                 `json:"directory,omitempty"`
	ContainerId         *string                                 `json:"containerId,omitempty"`
	Task                *SingularityTask                        `json:"task,omitempty"`
	HealthcheckResults  *SingularityTaskHealthcheckResultList   `json:"healthcheckResults,omitempty"`
	LoadBalancerUpdates *SingularityLoadBalancerUpdateList      `json:"loadBalancerUpdates,omitempty"`
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

func (self *SingularityTaskHistory) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskHistory) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskHistory) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHistory", name)

	case "shellCommandHistory", "ShellCommandHistory":
		v, ok := value.(SingularityTaskShellCommandHistoryList)
		if ok {
			self.ShellCommandHistory = &v
			return nil
		}
		return fmt.Errorf("Field shellCommandHistory/ShellCommandHistory: value %v(%T) couldn't be cast to type SingularityTaskShellCommandHistoryList", value, value)

	case "taskMetadata", "TaskMetadata":
		v, ok := value.(SingularityTaskMetadataList)
		if ok {
			self.TaskMetadata = &v
			return nil
		}
		return fmt.Errorf("Field taskMetadata/TaskMetadata: value %v(%T) couldn't be cast to type SingularityTaskMetadataList", value, value)

	case "taskUpdates", "TaskUpdates":
		v, ok := value.(SingularityTaskHistoryUpdateList)
		if ok {
			self.TaskUpdates = &v
			return nil
		}
		return fmt.Errorf("Field taskUpdates/TaskUpdates: value %v(%T) couldn't be cast to type SingularityTaskHistoryUpdateList", value, value)

	case "directory", "Directory":
		v, ok := value.(string)
		if ok {
			self.Directory = &v
			return nil
		}
		return fmt.Errorf("Field directory/Directory: value %v(%T) couldn't be cast to type string", value, value)

	case "containerId", "ContainerId":
		v, ok := value.(string)
		if ok {
			self.ContainerId = &v
			return nil
		}
		return fmt.Errorf("Field containerId/ContainerId: value %v(%T) couldn't be cast to type string", value, value)

	case "task", "Task":
		v, ok := value.(*SingularityTask)
		if ok {
			self.Task = v
			return nil
		}
		return fmt.Errorf("Field task/Task: value %v(%T) couldn't be cast to type *SingularityTask", value, value)

	case "healthcheckResults", "HealthcheckResults":
		v, ok := value.(SingularityTaskHealthcheckResultList)
		if ok {
			self.HealthcheckResults = &v
			return nil
		}
		return fmt.Errorf("Field healthcheckResults/HealthcheckResults: value %v(%T) couldn't be cast to type SingularityTaskHealthcheckResultList", value, value)

	case "loadBalancerUpdates", "LoadBalancerUpdates":
		v, ok := value.(SingularityLoadBalancerUpdateList)
		if ok {
			self.LoadBalancerUpdates = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerUpdates/LoadBalancerUpdates: value %v(%T) couldn't be cast to type SingularityLoadBalancerUpdateList", value, value)

	}
}

func (self *SingularityTaskHistory) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskHistory", name)

	case "shellCommandHistory", "ShellCommandHistory":
		return *self.ShellCommandHistory, nil
		return nil, fmt.Errorf("Field ShellCommandHistory no set on ShellCommandHistory %+v", self)

	case "taskMetadata", "TaskMetadata":
		return *self.TaskMetadata, nil
		return nil, fmt.Errorf("Field TaskMetadata no set on TaskMetadata %+v", self)

	case "taskUpdates", "TaskUpdates":
		return *self.TaskUpdates, nil
		return nil, fmt.Errorf("Field TaskUpdates no set on TaskUpdates %+v", self)

	case "directory", "Directory":
		return *self.Directory, nil
		return nil, fmt.Errorf("Field Directory no set on Directory %+v", self)

	case "containerId", "ContainerId":
		return *self.ContainerId, nil
		return nil, fmt.Errorf("Field ContainerId no set on ContainerId %+v", self)

	case "task", "Task":
		return self.Task, nil
		return nil, fmt.Errorf("Field Task no set on Task %+v", self)

	case "healthcheckResults", "HealthcheckResults":
		return *self.HealthcheckResults, nil
		return nil, fmt.Errorf("Field HealthcheckResults no set on HealthcheckResults %+v", self)

	case "loadBalancerUpdates", "LoadBalancerUpdates":
		return *self.LoadBalancerUpdates, nil
		return nil, fmt.Errorf("Field LoadBalancerUpdates no set on LoadBalancerUpdates %+v", self)

	}
}

func (self *SingularityTaskHistory) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHistory", name)

	case "shellCommandHistory", "ShellCommandHistory":
		self.ShellCommandHistory = nil

	case "taskMetadata", "TaskMetadata":
		self.TaskMetadata = nil

	case "taskUpdates", "TaskUpdates":
		self.TaskUpdates = nil

	case "directory", "Directory":
		self.Directory = nil

	case "containerId", "ContainerId":
		self.ContainerId = nil

	case "task", "Task":
		self.Task = nil

	case "healthcheckResults", "HealthcheckResults":
		self.HealthcheckResults = nil

	case "loadBalancerUpdates", "LoadBalancerUpdates":
		self.LoadBalancerUpdates = nil

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
