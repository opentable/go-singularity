package dtos

import "io"

type DeployProgress struct {
	AutoAdvanceDeploySteps bool
	DeployInstanceCountPerStep int32
	DeployStepWaitTimeMs int64
	FailedDeployTasks []
	StepComplete bool
	TargetActiveInstances int32
	Timestamp int64
}

func (self *DeployProgress) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DeployProgress) FormatText() string {
	return FormatText(self)
}

func (self *DeployProgress) FormatJSON() string {
	return FormatJSON(self)
}
