package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeployProgress struct {
	TargetActiveInstances      *int32                 `json:"targetActiveInstances,omitempty"`
	CurrentActiveInstances     *int32                 `json:"currentActiveInstances,omitempty"`
	DeployInstanceCountPerStep *int32                 `json:"deployInstanceCountPerStep,omitempty"`
	DeployStepWaitTimeMs       *int64                 `json:"deployStepWaitTimeMs,omitempty"`
	StepComplete               *bool                  `json:"stepComplete,omitempty"`
	AutoAdvanceDeploySteps     *bool                  `json:"autoAdvanceDeploySteps,omitempty"`
	FailedDeployTasks          *SingularityTaskIdList `json:"failedDeployTasks,omitempty"`
	Timestamp                  *int64                 `json:"timestamp,omitempty"`
}

func (self *SingularityDeployProgress) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployProgress) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployProgress); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployProgress cannot copy the values from %#v", other)
}

func (self *SingularityDeployProgress) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployProgress) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployProgress) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployProgress", name)

	case "targetActiveInstances", "TargetActiveInstances":
		v, ok := value.(int32)
		if ok {
			self.TargetActiveInstances = &v
			return nil
		}
		return fmt.Errorf("Field targetActiveInstances/TargetActiveInstances: value %v(%T) couldn't be cast to type int32", value, value)

	case "currentActiveInstances", "CurrentActiveInstances":
		v, ok := value.(int32)
		if ok {
			self.CurrentActiveInstances = &v
			return nil
		}
		return fmt.Errorf("Field currentActiveInstances/CurrentActiveInstances: value %v(%T) couldn't be cast to type int32", value, value)

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		v, ok := value.(int32)
		if ok {
			self.DeployInstanceCountPerStep = &v
			return nil
		}
		return fmt.Errorf("Field deployInstanceCountPerStep/DeployInstanceCountPerStep: value %v(%T) couldn't be cast to type int32", value, value)

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		v, ok := value.(int64)
		if ok {
			self.DeployStepWaitTimeMs = &v
			return nil
		}
		return fmt.Errorf("Field deployStepWaitTimeMs/DeployStepWaitTimeMs: value %v(%T) couldn't be cast to type int64", value, value)

	case "stepComplete", "StepComplete":
		v, ok := value.(bool)
		if ok {
			self.StepComplete = &v
			return nil
		}
		return fmt.Errorf("Field stepComplete/StepComplete: value %v(%T) couldn't be cast to type bool", value, value)

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		v, ok := value.(bool)
		if ok {
			self.AutoAdvanceDeploySteps = &v
			return nil
		}
		return fmt.Errorf("Field autoAdvanceDeploySteps/AutoAdvanceDeploySteps: value %v(%T) couldn't be cast to type bool", value, value)

	case "failedDeployTasks", "FailedDeployTasks":
		v, ok := value.(SingularityTaskIdList)
		if ok {
			self.FailedDeployTasks = &v
			return nil
		}
		return fmt.Errorf("Field failedDeployTasks/FailedDeployTasks: value %v(%T) couldn't be cast to type SingularityTaskIdList", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularityDeployProgress) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployProgress", name)

	case "targetActiveInstances", "TargetActiveInstances":
		return *self.TargetActiveInstances, nil
		return nil, fmt.Errorf("Field TargetActiveInstances no set on TargetActiveInstances %+v", self)

	case "currentActiveInstances", "CurrentActiveInstances":
		return *self.CurrentActiveInstances, nil
		return nil, fmt.Errorf("Field CurrentActiveInstances no set on CurrentActiveInstances %+v", self)

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		return *self.DeployInstanceCountPerStep, nil
		return nil, fmt.Errorf("Field DeployInstanceCountPerStep no set on DeployInstanceCountPerStep %+v", self)

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		return *self.DeployStepWaitTimeMs, nil
		return nil, fmt.Errorf("Field DeployStepWaitTimeMs no set on DeployStepWaitTimeMs %+v", self)

	case "stepComplete", "StepComplete":
		return *self.StepComplete, nil
		return nil, fmt.Errorf("Field StepComplete no set on StepComplete %+v", self)

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		return *self.AutoAdvanceDeploySteps, nil
		return nil, fmt.Errorf("Field AutoAdvanceDeploySteps no set on AutoAdvanceDeploySteps %+v", self)

	case "failedDeployTasks", "FailedDeployTasks":
		return *self.FailedDeployTasks, nil
		return nil, fmt.Errorf("Field FailedDeployTasks no set on FailedDeployTasks %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityDeployProgress) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployProgress", name)

	case "targetActiveInstances", "TargetActiveInstances":
		self.TargetActiveInstances = nil

	case "currentActiveInstances", "CurrentActiveInstances":
		self.CurrentActiveInstances = nil

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		self.DeployInstanceCountPerStep = nil

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		self.DeployStepWaitTimeMs = nil

	case "stepComplete", "StepComplete":
		self.StepComplete = nil

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		self.AutoAdvanceDeploySteps = nil

	case "failedDeployTasks", "FailedDeployTasks":
		self.FailedDeployTasks = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	}

	return nil
}

func (self *SingularityDeployProgress) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployProgressList []*SingularityDeployProgress

func (self *SingularityDeployProgressList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployProgressList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployProgressList cannot copy the values from %#v", other)
}

func (list *SingularityDeployProgressList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployProgressList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployProgressList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
