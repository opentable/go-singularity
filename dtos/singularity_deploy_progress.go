package dtos

import "io"

type SingularityDeployProgress struct {
	AutoAdvanceDeploySteps     bool                  `json:"autoAdvanceDeploySteps"`
	DeployInstanceCountPerStep int32                 `json:"deployInstanceCountPerStep"`
	DeployStepWaitTimeMs       int64                 `json:"deployStepWaitTimeMs"`
	FailedDeployTasks          SingularityTaskIdList `json:"failedDeployTasks"`
	StepComplete               bool                  `json:"stepComplete"`
	TargetActiveInstances      int32                 `json:"targetActiveInstances"`
	Timestamp                  int64                 `json:"timestamp"`
}

func (self *SingularityDeployProgress) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployProgress) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployProgress) FormatJSON() string {
	return FormatJSON(self)
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
