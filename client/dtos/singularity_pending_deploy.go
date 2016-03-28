package dtos

import "io"

type SingularityPendingDeploy struct {
	//	CurrentDeployState *DeployState
	DeployMarker           *SingularityDeployMarker
	DeployProgress         *SingularityDeployProgress
	LastLoadBalancerUpdate *SingularityLoadBalancerUpdate
}

func (self *SingularityPendingDeploy) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPendingDeploy) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPendingDeploy) FormatJSON() string {
	return FormatJSON(self)
}
