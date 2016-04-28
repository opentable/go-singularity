package dtos

import (
	"fmt"
	"io"
)

type SingularityDeployProgress struct {
	present                    map[string]bool
	AutoAdvanceDeploySteps     bool                  `json:"autoAdvanceDeploySteps"`
	DeployInstanceCountPerStep int32                 `json:"deployInstanceCountPerStep"`
	DeployStepWaitTimeMs       int64                 `json:"deployStepWaitTimeMs"`
	FailedDeployTasks          SingularityTaskIdList `json:"failedDeployTasks"`
	StepComplete               bool                  `json:"stepComplete"`
	TargetActiveInstances      int32                 `json:"targetActiveInstances"`
	Timestamp                  int64                 `json:"timestamp"`
}

func (self *SingularityDeployProgress) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployProgress) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *SingularityDeployProgress) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployProgress) FormatJSON() string {
	return FormatJSON(self)
}

func (self *SingularityDeployProgress) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *SingularityDeployProgress) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployProgress", name)

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		v, ok := value.(bool)
		if ok {
			self.AutoAdvanceDeploySteps = v
			self.present["autoAdvanceDeploySteps"] = true
			return nil
		} else {
			return fmt.Errorf("Field autoAdvanceDeploySteps/AutoAdvanceDeploySteps: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		v, ok := value.(int32)
		if ok {
			self.DeployInstanceCountPerStep = v
			self.present["deployInstanceCountPerStep"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployInstanceCountPerStep/DeployInstanceCountPerStep: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		v, ok := value.(int64)
		if ok {
			self.DeployStepWaitTimeMs = v
			self.present["deployStepWaitTimeMs"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployStepWaitTimeMs/DeployStepWaitTimeMs: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "failedDeployTasks", "FailedDeployTasks":
		v, ok := value.(SingularityTaskIdList)
		if ok {
			self.FailedDeployTasks = v
			self.present["failedDeployTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field failedDeployTasks/FailedDeployTasks: value %v(%T) couldn't be cast to type SingularityTaskIdList", value, value)
		}

	case "stepComplete", "StepComplete":
		v, ok := value.(bool)
		if ok {
			self.StepComplete = v
			self.present["stepComplete"] = true
			return nil
		} else {
			return fmt.Errorf("Field stepComplete/StepComplete: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "targetActiveInstances", "TargetActiveInstances":
		v, ok := value.(int32)
		if ok {
			self.TargetActiveInstances = v
			self.present["targetActiveInstances"] = true
			return nil
		} else {
			return fmt.Errorf("Field targetActiveInstances/TargetActiveInstances: value %v(%T) couldn't be cast to type int32", value, value)
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

func (self *SingularityDeployProgress) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployProgress", name)

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		if self.present != nil {
			if _, ok := self.present["autoAdvanceDeploySteps"]; ok {
				return self.AutoAdvanceDeploySteps, nil
			}
		}
		return nil, fmt.Errorf("Field AutoAdvanceDeploySteps no set on AutoAdvanceDeploySteps %+v", self)

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		if self.present != nil {
			if _, ok := self.present["deployInstanceCountPerStep"]; ok {
				return self.DeployInstanceCountPerStep, nil
			}
		}
		return nil, fmt.Errorf("Field DeployInstanceCountPerStep no set on DeployInstanceCountPerStep %+v", self)

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		if self.present != nil {
			if _, ok := self.present["deployStepWaitTimeMs"]; ok {
				return self.DeployStepWaitTimeMs, nil
			}
		}
		return nil, fmt.Errorf("Field DeployStepWaitTimeMs no set on DeployStepWaitTimeMs %+v", self)

	case "failedDeployTasks", "FailedDeployTasks":
		if self.present != nil {
			if _, ok := self.present["failedDeployTasks"]; ok {
				return self.FailedDeployTasks, nil
			}
		}
		return nil, fmt.Errorf("Field FailedDeployTasks no set on FailedDeployTasks %+v", self)

	case "stepComplete", "StepComplete":
		if self.present != nil {
			if _, ok := self.present["stepComplete"]; ok {
				return self.StepComplete, nil
			}
		}
		return nil, fmt.Errorf("Field StepComplete no set on StepComplete %+v", self)

	case "targetActiveInstances", "TargetActiveInstances":
		if self.present != nil {
			if _, ok := self.present["targetActiveInstances"]; ok {
				return self.TargetActiveInstances, nil
			}
		}
		return nil, fmt.Errorf("Field TargetActiveInstances no set on TargetActiveInstances %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityDeployProgress) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployProgress", name)

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		self.present["autoAdvanceDeploySteps"] = false

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		self.present["deployInstanceCountPerStep"] = false

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		self.present["deployStepWaitTimeMs"] = false

	case "failedDeployTasks", "FailedDeployTasks":
		self.present["failedDeployTasks"] = false

	case "stepComplete", "StepComplete":
		self.present["stepComplete"] = false

	case "targetActiveInstances", "TargetActiveInstances":
		self.present["targetActiveInstances"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	}

	return nil
}

func (self *SingularityDeployProgress) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type SingularityDeployProgressList []*SingularityDeployProgress

func (list *SingularityDeployProgressList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
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
	return FormatJSON(list)
}
