package dtos

import "io"

type SingularityDeployProgress struct {
	AutoAdvanceDeploySteps     bool
	DeployInstanceCountPerStep int32
	DeployStepWaitTimeMs       int64
	FailedDeployTasks          *SingularityTaskId
	StepComplete               bool
	TargetActiveInstances      int32
	Timestamp                  int64
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
