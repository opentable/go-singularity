package dtos

import "io"

type SingularityDeployResult struct {
	DeployFailures *SingularityDeployFailure
	//	DeployState *DeployState
	LbUpdate  *SingularityLoadBalancerUpdate
	Message   string
	Timestamp int64
}

func (self *SingularityDeployResult) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployResult) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployResult) FormatJSON() string {
	return FormatJSON(self)
}
